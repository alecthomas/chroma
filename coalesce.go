package chroma

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer {
	return &coalescer{lexer}
}

type coalescer struct {
	Lexer
}

func (d *coalescer) Tokenise(text string) ([]Token, error) {
	in, err := d.Lexer.Tokenise(text)
	if err != nil {
		return in, err
	}
	out := []Token{}
	for _, token := range in {
		if len(out) == 0 {
			out = append(out, token)
			continue
		}
		last := &out[len(out)-1]
		if last.Type == token.Type {
			last.Value += token.Value
		} else {
			out = append(out, token)
		}
	}
	return out, err
}
