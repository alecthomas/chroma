package formatters

import (
	"io"

	"github.com/alecthomas/chroma"
)

// Formatter takes a token stream and formats it.
type Formatter interface {
	Format(w io.Writer, tokens []chroma.Token) error
}
