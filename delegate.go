package chroma

import (
	"bytes"
	"iter"
	"slices"
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

func (d *delegatingLexer) SetTracing(enable bool) {
	if l, ok := d.language.(TracingLexer); ok {
		l.SetTracing(enable)
	}
	if l, ok := d.root.(TracingLexer); ok {
		l.SetTracing(enable)
	}
}

func (d *delegatingLexer) AnalyseText(text string) float32 {
	return d.root.AnalyseText(text)
}

func (d *delegatingLexer) SetAnalyser(analyser func(text string) float32) Lexer {
	d.root.SetAnalyser(analyser)
	return d
}

func (d *delegatingLexer) SetRegistry(r *LexerRegistry) Lexer {
	d.root.SetRegistry(r)
	d.language.SetRegistry(r)
	return d
}

func (d *delegatingLexer) Config() *Config {
	return d.language.Config()
}

// An insertion is the character range where language tokens should be inserted.
type insertion struct {
	start, end int
	tokens     []Token
}

func (d *delegatingLexer) Tokenise(options *TokeniseOptions, text string) (iter.Seq[Token], error) { // nolint: gocognit
	tokens, err := Tokenise(Coalesce(d.language), options, text)
	if err != nil {
		return nil, err
	}
	// Compute insertions and gather "Other" tokens.
	others := &bytes.Buffer{}
	insertions := []*insertion{}
	var insert *insertion
	offset := 0
	first := true
	var lastType TokenType
	for _, t := range tokens {
		if t.Type == Other {
			if !first && insert != nil && lastType != Other {
				insert.end = offset
			}
			others.WriteString(t.Value)
		} else {
			if first || lastType == Other {
				insert = &insertion{start: offset}
				insertions = append(insertions, insert)
			}
			insert.tokens = append(insert.tokens, t)
		}
		first = false
		lastType = t.Type
		offset += len(t.Value)
	}

	if len(insertions) == 0 {
		return d.root.Tokenise(options, text)
	}

	// Lex the other tokens.
	rootTokens, err := Tokenise(Coalesce(d.root), options, others.String())
	if err != nil {
		return nil, err
	}

	// Interleave the two sets of tokens.
	var out []Token
	offset = 0
	ti := 0
	ii := 0
	for ti < len(rootTokens) || ii < len(insertions) {
		if ti >= len(rootTokens) || (ii < len(insertions) && insertions[ii].start < offset+len(rootTokens[ti].Value)) {
			ins := insertions[ii]
			ii++
			if ti < len(rootTokens) {
				l, r := splitToken(rootTokens[ti], ins.start-offset)
				if l.Value != "" {
					out = append(out, l)
					offset += len(l.Value)
				}
				rootTokens[ti] = r
			}
			out = append(out, ins.tokens...)
			offset += ins.end - ins.start
		} else {
			out = append(out, rootTokens[ti])
			offset += len(rootTokens[ti].Value)
			ti++
		}
	}
	return slices.Values(out), nil
}

func splitToken(t Token, offset int) (l Token, r Token) {
	if offset <= 0 {
		return Token{}, t
	}
	if offset >= len(t.Value) {
		return t, Token{}
	}
	l = t.Clone()
	r = t.Clone()
	l.Value = l.Value[:offset]
	r.Value = r.Value[offset:]
	return
}
