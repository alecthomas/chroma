package formatters

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/styles"
)

// TTY16m is a true-colour terminal formatter.
var TTY16m = Register("terminal16m", FormatterFunc(trueColourFormatter))

func trueColourFormatter(w io.Writer, style *styles.Style) (func(*chroma.Token), error) {
	return func(token *chroma.Token) {
		entry := style.Get(token.Type)
		if !entry.IsZero() {
			out := ""
			if entry.Bold {
				out += "\033[1m"
			}
			if entry.Underline {
				out += "\033[4m"
			}
			if entry.Colour.IsSet() {
				out += fmt.Sprintf("\033[38:2:%d:%d:%dm", entry.Colour.Red(), entry.Colour.Green(), entry.Colour.Blue())
			}
			if entry.Background.IsSet() {
				out += fmt.Sprintf("\033[48:2:%d:%d:%dm", entry.Background.Red(), entry.Background.Green(), entry.Background.Blue())
			}
			fmt.Fprint(w, out)
		}
		fmt.Fprint(w, token.Value)
		if !entry.IsZero() {
			fmt.Fprint(w, "\033[0m")
		}
	}, nil
}
