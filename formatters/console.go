package formatters

import (
	"bufio"
	"io"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/colour"
)

var DefaultConsoleTheme = map[TokenType]string{
	Number:  "^B^3",
	Comment: "^5",
	String:  "^B^5",
	Keyword: "^B^7",
}

// Console formatter.
//
// 		formatter := Console(DefaultConsoleTheme)
func Console(theme map[TokenType]string) Formatter {
	return &consoleFormatter{theme}
}

type consoleFormatter struct {
	theme map[TokenType]string
}

func (c *consoleFormatter) Format(w io.Writer, tokens []Token) error {
	bw := bufio.NewWriterSize(w, 1024)
	printer := colour.Colour(bw)
	for _, token := range tokens {
		clr, ok := c.theme[token.Type]
		if !ok {
			clr, ok = c.theme[token.Type.SubCategory()]
			if !ok {
				clr, ok = c.theme[token.Type.Category()]
				if !ok {
					clr = "^R"
				}
			}
		}
		printer.Printf(clr+"%s", token.Value)
	}
	bw.Flush()
	return nil
}
