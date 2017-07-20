package lexers

import (
	"testing"

	"github.com/alecthomas/chroma"
)

const lexerBenchSource = `package chroma

import (
	"io"
)

// A Formatter for Chroma lexers.
type Formatter interface {
	// Format returns a formatting function for tokens.
	Format(w io.Writer, style *Style) (func(*Token), error)
}

// A FormatterFunc is a Formatter implemented as a function.
type FormatterFunc func(io.Writer, *Style) (func(*Token), error)

func (f FormatterFunc) Format(w io.Writer, s *Style) (func(*Token), error) {
	return f(w, s)
}
`

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Go.Tokenise(nil, lexerBenchSource, func(t *chroma.Token) {})
	}
}
