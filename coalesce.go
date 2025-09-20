package chroma

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer { return &coalescer{lexer} }

type coalescer struct{ Lexer }

func (d *coalescer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	it, err := d.Lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	return func(yield func(Token) bool) {
		var prev Token
		for token := range it {
			if len(token.Value) == 0 {
				continue
			}
			if prev.IsZero() {
				prev = token
			} else {
				if prev.Type == token.Type && len(prev.Value) < 8192 {
					prev.Value += token.Value
				} else {
					if !yield(prev) {
						return
					}
					prev = token
				}
			}
		}
		if !prev.IsZero() {
			yield(prev)
		}
	}, nil
}
