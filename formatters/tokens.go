package formatters

import (
	"fmt"
	"io"
	"iter"

	"github.com/alecthomas/chroma/v2"
)

// Tokens formatter outputs the raw token structures.
var Tokens = Register("tokens", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style, it iter.Seq[chroma.Token]) error {
	for t := range it {
		if t == chroma.EOF {
			break
		}
		if _, err := fmt.Fprintln(w, t.GoString()); err != nil {
			return err
		}
	}
	return nil
}))
