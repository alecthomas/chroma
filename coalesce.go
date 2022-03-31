package chroma

import (
	"fmt"
)

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer { return &coalescer{lexer} }

type coalescer struct{ Lexer }

func (d *coalescer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	it, err := d.Lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	return d.iter(it), nil
}

func (d *coalescer) iter(it func() Token) func() Token {
	var prev Token
	return func() Token {
		for token := it(); token != (EOF); token = it() {
			if len(token.Value) == 0 {
				continue
			}
			if prev == EOF {
				prev = token
			} else {
				if prev.Type == token.Type && len(prev.Value) < 8192 {
					prev.Value += token.Value
				} else {
					out := prev
					prev = token
					return out
				}
			}
		}
		out := prev
		prev = EOF
		return out
	}
}

func (d *coalescer) TokeniseWithOriginalLen(options *TokeniseOptions, text string) (Iterator, OriginalLenIterator, error) {
	lex, ok := d.Lexer.(TokeniserWithOriginalLen)

	if !ok {
		err := fmt.Errorf("lexer does not support tokenizing with offsets")
		return nil, OriginalLenIterator{}, err
	}

	it, offsetIter, err := lex.TokeniseWithOriginalLen(options, text)
	if err != nil {
		return nil, OriginalLenIterator{}, err
	}

	return d.iter(it), offsetIter, nil
}
