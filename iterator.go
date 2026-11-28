package chroma

import (
	"iter"
	"slices"
	"strings"
)

// Concaterator concatenates tokens from a series of iterators.
func Concaterator(iterators ...iter.Seq[Token]) iter.Seq[Token] {
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

// Literator converts a sequence of literal Tokens into an iterator.
func Literator(tokens ...Token) iter.Seq[Token] {
	return slices.Values(tokens)
}

// SplitTokensIntoLines lazily splits tokens into lines, splitting tokens containing newlines in two.
func SplitTokensIntoLines(tokens iter.Seq[Token]) iter.Seq[[]Token] {
	return func(yield func([]Token) bool) {
		var line []Token
	tokenLoop:
		for token := range tokens {
			for strings.Contains(token.Value, "\n") {
				parts := strings.SplitAfterN(token.Value, "\n", 2)
				// Token becomes the tail.
				token.Value = parts[1]

				// Append the head to the line and flush the line.
				if !yield(append(line, Token{token.Type, parts[0]})) {
					return
				}
				line = nil

				// If the tail token is empty, don't emit it.
				if len(token.Value) == 0 {
					continue tokenLoop
				}
			}
			line = append(line, token)
		}
		// Strip empty trailing token line.
		if len(line) == 0 || (len(line) == 1 && line[0].Value == "") {
			return
		}
		yield(line)
	}
}
