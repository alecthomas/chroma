package formatters

import (
	"fmt"
	"io"
	"regexp"

	"github.com/alecthomas/chroma/v2"
)

// TTY16m is a true-colour terminal formatter.
var TTY16m = Register("terminal16m", chroma.FormatterFunc(trueColourFormatter))

var crOrCrLf = regexp.MustCompile(`\r?\n`)

func trueColorTokenFormatter(w io.Writer, formatting string, text string) {
	fmt.Fprint(w, formatting)
	fmt.Fprint(w, text)
	fmt.Fprint(w, "\033[0m")
}

func trueColourFormatter(w io.Writer, style *chroma.Style, it chroma.Iterator) error {
	style = clearBackground(style)
	for token := it(); token != chroma.EOF; token = it() {
		entry := style.Get(token.Type)
		if entry.IsZero() {
			fmt.Fprint(w, token.Value)
			continue
		}

		formatting := ""
		if entry.Bold == chroma.Yes {
			formatting += "\033[1m"
		}
		if entry.Underline == chroma.Yes {
			formatting += "\033[4m"
		}
		if entry.Italic == chroma.Yes {
			formatting += "\033[3m"
		}
		if entry.Colour.IsSet() {
			formatting += fmt.Sprintf("\033[38;2;%d;%d;%dm", entry.Colour.Red(), entry.Colour.Green(), entry.Colour.Blue())
		}
		if entry.Background.IsSet() {
			formatting += fmt.Sprintf("\033[48;2;%d;%d;%dm", entry.Background.Red(), entry.Background.Green(), entry.Background.Blue())
		}

		trueColorTokenFormatter(w, formatting, token.Value)
	}
	return nil
}
