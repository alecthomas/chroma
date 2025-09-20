package chroma

import (
	"iter"
	"slices"
	"strings"
)

// An Iterator across tokens.
//
// EOF will be returned at the end of the Token stream.
//
// If an error occurs within an Iterator, it may propagate this in a panic. Formatters should recover.
type Iterator = iter.Seq[Token]

// Concaterator concatenates tokens from a series of iterators.
func Concaterator(iterators ...Iterator) Iterator {
	return func(yield func(Token) bool) {
		for _, it := range iterators {
			for t := range it {
				if !yield(t) {
					return
				}
			}
		}
	}
}

// Literator converts a sequence of literal Tokens into an Iterator.
func Literator(tokens ...Token) Iterator {
	return slices.Values(tokens)
}

// SplitTokensIntoLines splits tokens containing newlines in two.
func SplitTokensIntoLines(tokens Iterator) iter.Seq[[]Token] {
	return func(yield func([]Token) bool) {
		var line []Token
		for token := range tokens {
			line = slices.Grow(line, strings.Count(token.Value, "\n")+1)
			for part := range strings.SplitAfterSeq(token.Value, "\n") {
				if part == "" {
					continue // Empty substring at end of token. Ignore
				}
				token := token
				token.Value = part
				line = append(line, token)
				if strings.HasSuffix(part, "\n") {
					if !yield(line) {
						return
					}
					line = nil
				}
			}
		}
		if len(line) > 0 {
			yield(line)
		}
	}
}
