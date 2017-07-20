package chroma

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer {
	return &coalescer{lexer}
}

type coalescer struct {
	Lexer
}

func (d *coalescer) Tokenise(options *TokeniseOptions, text string, out func(*Token)) error {
	var prev *Token
	return d.Lexer.Tokenise(options, text, func(token *Token) {
		if prev == nil {
			prev = token
		} else {
			if prev.Type == token.Type && len(prev.Value) < 8192 {
				prev.Value += token.Value
			} else {
				out(prev)
				prev = token
			}
		}
		if token.Type == EOF {
			out(token)
		}
	})
}
