package formatters

import (
	"encoding/json"
	"fmt"
	"io"
	"iter"

	"github.com/alecthomas/chroma/v2"
)

// JSON formatter outputs the raw token structures as JSON.
var JSON = Register("json", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style, it iter.Seq[chroma.Token]) error {
	if _, err := fmt.Fprintln(w, "["); err != nil {
		return err
	}
	i := 0
	for t := range it {
		if t == chroma.EOF {
			break
		}
		if i > 0 {
			if _, err := fmt.Fprintln(w, ","); err != nil {
				return err
			}
		}
		i++
		bytes, err := json.Marshal(t)
		if err != nil {
			return err
		}
		if _, err := fmt.Fprint(w, "  "+string(bytes)); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "]"); err != nil {
		return err
	}
	return nil
}))
