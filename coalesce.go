package chroma

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer { return &coalescer{lexer} }

type coalescer struct{ Lexer }

func (d *coalescer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	var prev Token
	havePrev := false
	it, err := d.Lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	return func() (Token, bool) {
		for {
			token, ok := it()
			if !ok {
				break
			}
			if len(token.Value) == 0 {
				continue
			}
			if !havePrev {
				prev = token
				havePrev = true
			} else {
				if prev.Type == token.Type && len(prev.Value) < 8192 {
					prev.Value += token.Value
				} else {
					out := prev
					prev = token
					return out, true
				}
			}
		}
		if havePrev {
			havePrev = false
			return prev, true
		}
		return Token{}, false
	}, nil
}
