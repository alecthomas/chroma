package chroma

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer {
	return &coalescer{lexer}
}

type coalescer struct {
	Lexer
}

func (d *coalescer) Tokenise(options *TokeniseOptions, text string, out func(Token)) error {
	var last *Token
	defer func() {
		if last != nil {
			out(*last)
		}
	}()
	return d.Lexer.Tokenise(options, text, func(token Token) {
		if last == nil {
			last = &token
		} else {
			if last.Type == token.Type {
				last.Value += token.Value
			} else {
				out(*last)
				last = &token
			}
		}
	})
}
