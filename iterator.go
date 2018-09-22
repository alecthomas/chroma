package chroma

// An Iterator across tokens.
//
// Token{}, false will be returned at the end of the Token stream.
//
// If an error occurs within an Iterator, it may propagate this in a panic. Formatters should recover.
type Iterator func() (Token, bool)

// Tokens consumes all tokens from the iterator and returns them as a slice.
func (i Iterator) Tokens() []Token {
	out := []Token{}
	for {
		t, ok := i()
		if !ok {
			break
		}
		out = append(out, t)
	}
	return out
}

// Concaterator concatenates tokens from a series of iterators.
func Concaterator(iterators ...Iterator) Iterator {
	return func() (Token, bool) {
		for len(iterators) > 0 {
			t, ok := iterators[0]()
			if ok {
				return t, ok
			}
			iterators = iterators[1:]
		}
		return Token{}, false
	}
}

// Literator converts a sequence of literal Tokens into an Iterator.
func Literator(tokens ...Token) Iterator {
	return func() (Token, bool) {
		if len(tokens) == 0 {
			return Token{}, false
		}
		token := tokens[0]
		tokens = tokens[1:]
		return token, true
	}
}
