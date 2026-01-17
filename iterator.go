package chroma

import (
	"iter"
	"strings"
)

// Concaterator concatenates tokens from a series of iterators.
func Concaterator(iterators ...iter.Seq[Token]) iter.Seq[Token] {
	return func(yield func(Token) bool) {
		for _, it := range iterators {
			for t := range it {
				if t == EOF {
					break
				}
				if !yield(t) {
					return
				}
			}
		}
	}
}

// Literator converts a sequence of literal Tokens into an iter.Seq[Token].
func Literator(tokens ...Token) iter.Seq[Token] {
	return func(yield func(Token) bool) {
		for _, token := range tokens {
			if !yield(token) {
				return
			}
		}
	}
}

// SplitTokensIntoLines splits tokens containing newlines in two.
func SplitTokensIntoLines(tokens []Token) (out [][]Token) {
	var line []Token // nolint: prealloc
tokenLoop:
	for _, token := range tokens {
		for strings.Contains(token.Value, "\n") {
			parts := strings.SplitAfterN(token.Value, "\n", 2)
			// Token becomes the tail.
			token.Value = parts[1]

			// Append the head to the line and flush the line.
			clone := token.Clone()
			clone.Value = parts[0]
			line = append(line, clone)
			out = append(out, line)
			line = nil

			// If the tail token is empty, don't emit it.
			if len(token.Value) == 0 {
				continue tokenLoop
			}
		}
		line = append(line, token)
	}
	if len(line) > 0 {
		out = append(out, line)
	}
	// Strip empty trailing token line.
	if len(out) > 0 {
		last := out[len(out)-1]
		if len(last) == 1 && last[0].Value == "" {
			out = out[:len(out)-1]
		}
	}
	return out
}
