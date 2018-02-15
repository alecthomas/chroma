package lexers_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma/lexers/g"
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
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		it, err := g.Go.Tokenise(nil, lexerBenchSource)
		assert.NoError(b, err)
		for t := it(); t != nil; t = it() {
		}
	}
}
