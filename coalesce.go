package chroma

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer, keepLineSplits bool) Lexer { return &coalescer{lexer, keepLineSplits} }

type coalescer struct {
	Lexer
	keepLineSplits bool
}

func (d *coalescer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	var prev Token
	it, err := d.Lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	return func() Token {
		for token := it(); token != (EOF); token = it() {
			if len(token.Value) == 0 {
				continue
			}
			if prev == EOF {
				prev = token
			} else {
				if prev.Type == token.Type && len(prev.Value) < 8192 && (!d.keepLineSplits || prev.Value[len(prev.Value)-1] != '\n') {
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
	}, nil
}
