package chroma

import "iter"

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer { return &coalescer{lexer} }

type coalescer struct{ Lexer }

func (d *coalescer) SetAnalyser(analyser func(text string) float32) Lexer {
	d.Lexer.SetAnalyser(analyser)
	return d
}

func (d *coalescer) SetRegistry(registry *LexerRegistry) Lexer {
	d.Lexer.SetRegistry(registry)
	return d
}

func (d *coalescer) SetTracing(enable bool) {
	if l, ok := d.Lexer.(TracingLexer); ok {
		l.SetTracing(enable)
	}
}

func (d *coalescer) Tokenise(options *TokeniseOptions, text string) (iter.Seq[Token], error) {
	it, err := d.Lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	return func(yield func(Token) bool) {
		var prev *Token
		for token := range it {
			if len(token.Value) == 0 {
				continue
			}
			if prev == nil {
				t := token
				prev = &t
			} else if prev.Type == token.Type && len(prev.Value) < 8192 {
				prev.Value += token.Value
			} else {
				if !yield(*prev) {
					return
				}
				t := token
				prev = &t
			}
		}
		if prev != nil {
			yield(*prev)
		}
	}, nil
}
