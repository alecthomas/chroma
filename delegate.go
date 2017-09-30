package chroma

import (
	"bytes"
)

type delegatingLexer struct {
	root     Lexer
	language Lexer
}

// DelegatingLexer takes two lexer as arguments. A root lexer and
// a language lexer. First everything is scanned using the language
// lexer, afterwards all Other tokens are lexed using the root
// lexer.
//
// The lexers from the template lexer package use this base lexer.
func DelegatingLexer(root Lexer, language Lexer) Lexer {
	return &delegatingLexer{
		root:     root,
		language: language,
	}
}

func (d *delegatingLexer) Config() *Config {
	return d.language.Config()
}

type tokenSplit struct {
	pos    int
	tokens []*Token
}

func splitOtherTokens(it Iterator) ([]tokenSplit, string) {
	splits := []tokenSplit{}
	var split *tokenSplit
	other := bytes.Buffer{}
	offset := 0
	for t := it(); t != nil; t = it() {
		if t.Type == Other {
			if split != nil {
				splits = append(splits, *split)
				split = nil
			}
			other.WriteString(t.Value)
		} else {
			if split == nil {
				split = &tokenSplit{pos: offset}
			}
			split.tokens = append(split.tokens, t)
		}
		offset += len(t.Value)
	}
	if split != nil {
		splits = append(splits, *split)
	}
	return splits, other.String()
}

func (d *delegatingLexer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	it, err := d.language.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	splits, other := splitOtherTokens(it)
	it, err = d.root.Tokenise(options, other)
	if err != nil {
		return nil, err
	}

	offset := 0
	return func() *Token {
		// First, see if there's a split at the start of this token.
		for len(splits) > 0 && splits[0].pos == offset {
			if len(splits[0].tokens) > 0 {
				t := splits[0].tokens[0]
				splits[0].tokens = splits[0].tokens[1:]
				offset += len(t.Value)
				return t
			}
			// End of tokens from this split, shift it off the queue.
			splits = splits[1:]
		}

		// No split, try to consume a token.
		t := it()
		if t == nil {
			for len(splits) > 0 {
				if len(splits[0].tokens) > 0 {
					t = splits[0].tokens[0]
					splits[0].tokens = splits[0].tokens[1:]
					offset += len(t.Value)
					return t
				}
				// End of tokens from this split, shift it off the queue.
				splits = splits[1:]
			}
			return nil
		}

		// Check if there's a split in the middle of the current token.
		if len(splits) > 0 && splits[0].pos < offset+len(t.Value) {
			// Split the token.
			next := t.Clone()
			point := splits[0].pos - offset
			next.Value = next.Value[point:]
			t.Value = t.Value[:point]

			// Insert the tail of the split token after any other splits at the same point.
			tailPos := offset + len(t.Value)
			tail := []tokenSplit{{pos: tailPos, tokens: []*Token{next}}}
			i := 0
			for ; i < len(splits); i++ {
				if splits[i].pos > tailPos {
					break
				}
			}
			splits = append(splits[:i], append(tail, splits[i:]...)...)

			// Finally, return the head.
		}

		offset += len(t.Value)
		return t
	}, nil
}
