package chroma

// An Iterator across tokens.
//
// nil will be returned at the end of the Token stream.
type Iterator func() *Token

// Concaterator concatenates tokens from a series of iterators.
func Concaterator(iterators ...Iterator) Iterator {
	return func() *Token {
		for len(iterators) > 0 {
			t := iterators[0]()
			if t != nil {
				return t
			}
			iterators = iterators[1:]
		}
		return nil
	}
}

// Literator converts a sequence of literal Tokens into an Iterator.
func Literator(tokens ...*Token) Iterator {
	return func() (out *Token) {
		if len(tokens) == 0 {
			return nil
		}
		token := tokens[0]
		tokens = tokens[1:]
		return token
	}
}

// Flatten an Iterator into its tokens.
func Flatten(iterator Iterator) []*Token {
	out := []*Token{}
	for t := iterator(); t != nil; t = iterator() {
		out = append(out, t)
	}
	return out
}
