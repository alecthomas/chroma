package formatters

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma"
)

// Raw formatter outputs the raw token structures.
var Raw = Register("raw", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style) (func(*chroma.Token), error) {
	return func(token *chroma.Token) {
		fmt.Fprintln(w, token.GoString())
	}, nil
}))
