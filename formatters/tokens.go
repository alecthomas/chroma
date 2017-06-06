package formatters

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/styles"
)

// Tokens formatter outputs the raw token structures.
var Tokens = Register("tokens", FormatterFunc(func(w io.Writer, s *styles.Style) (func(*chroma.Token), error) {
	return func(token *chroma.Token) {
		fmt.Fprintln(w, token.GoString())
	}, nil
}))
