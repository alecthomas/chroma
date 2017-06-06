package formatters

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma"
)

// Tokens formatter outputs the raw token structures.
var Tokens = Register("tokens", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style) (func(*chroma.Token), error) {
	return func(token *chroma.Token) {
		fmt.Fprintln(w, token.GoString())
	}, nil
}))
