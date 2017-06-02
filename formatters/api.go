package formatters

import (
	"io"

	"github.com/alecthomas/chroma"
)

// Formatter returns a formatting function for tokens.
type Formatter interface {
	Format(w io.Writer) (func(chroma.Token), error)
}
