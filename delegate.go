package chroma

import (
	"bytes"
)

type delegatingLexer struct {
	root     Lexer
	language Lexer
}

// DelegatingLexer combines two lexers to handle the common case of a language embedded inside another, such as PHP
// inside HTML or PHP inside plain text.
//
// It takes two lexer as arguments: a root lexer and a language lexer.  First everything is scanned using the language
// lexer, which must return "Other" for unrecognised tokens. Then all "Other" tokens are lexed using the root lexer.
// Finally, these two sets of tokens are merged.
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

// An insertion is the character range where language tokens should be inserted.
type insertion struct {
	start, end int
	tokens     []*Token
}

func (d *delegatingLexer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	tokens, err := Tokenise(Coalesce(d.language), options, text)
	if err != nil {
		return nil, err
	}
	// Compute insertions and gather "Other" tokens.
	others := &bytes.Buffer{}
	insertions := []*insertion{}
	var insert *insertion
	offset := 0
	var last *Token
	for _, t := range tokens {
		if t.Type == Other {
			if last != nil && insert != nil && last.Type != Other {
				insert.end = offset
			}
			others.WriteString(t.Value)
		} else {
			if last == nil || last.Type == Other {
				insert = &insertion{start: offset}
				insertions = append(insertions, insert)
			}
			insert.tokens = append(insert.tokens, t)
		}
		last = t
		offset += len(t.Value)
	}

	if len(insertions) == 0 {
		return d.root.Tokenise(options, text)
	}

	// Lex the other tokens.
	rootTokens, err := Tokenise(d.root, options, others.String())
	if err != nil {
		return nil, err
	}

	// Interleave the two sets of tokens.
	out := []*Token{}
	offset = 0
	index := 0
	next := func() *Token {
		if index >= len(rootTokens) {
			return nil
		}
		t := rootTokens[index]
		index++
		return t
	}
	t := next()
	for _, insert := range insertions {
		// Consume tokens until insertion point.
		for t != nil && offset+len(t.Value) <= insert.start {
			out = append(out, t)
			offset += len(t.Value)
			t = next()
		}
		// End of root tokens, append insertion point.
		if t == nil {
			out = append(out, insert.tokens...)
			break
		}

		// Split and insert.
		l, r := splitToken(t, insert.start-offset)
		if l != nil {
			out = append(out, l)
			offset += len(l.Value)
		}
		out = append(out, insert.tokens...)
		offset += insert.end - insert.start
		if r != nil {
			out = append(out, r)
			offset += len(r.Value)
		}
		t = next()
	}
	if t != nil {
		out = append(out, t)
	}
	// Remainder.
	out = append(out, rootTokens[index:]...)
	return Literator(out...), nil
}

func splitToken(t *Token, offset int) (l *Token, r *Token) {
	if offset == 0 {
		return nil, t
	}
	if offset >= len(t.Value) {
		return t, nil
	}
	l = t.Clone()
	r = t.Clone()
	l.Value = l.Value[:offset]
	r.Value = r.Value[offset:]
	return
}
