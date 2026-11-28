package chroma

import "iter"

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer { return &coalescer{lexer: lexer} }

type coalescer struct{ lexer Lexer }

func (d *coalescer) Config() *Config {
	return d.lexer.Config()
}

func (d *coalescer) AnalyseText(text string) float32 {
	return d.lexer.AnalyseText(text)
}

func (d *coalescer) SetAnalyser(analyser func(text string) float32) Lexer {
	d.lexer.SetAnalyser(analyser)
	return d
}

func (d *coalescer) SetRegistry(registry *LexerRegistry) Lexer {
	d.lexer.SetRegistry(registry)
	return d
}

func (d *coalescer) Tokenise(text string, options ...TokeniseOption) (iter.Seq[Token], error) {
	it, err := d.lexer.Tokenise(text, options...)
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
