package formatters

import (
	"fmt"
	"io"

	. "github.com/alecthomas/chroma" // nolint
)

var DefaultConsoleTheme = map[TokenType]string{
	Number:            "\033[1m\033[33m",
	Comment:           "\033[36m",
	String:            "\033[1m\033[36m",
	Keyword:           "\033[1m\033[37m",
	GenericHeading:    "\033[1m",
	GenericSubheading: "\033[1m",
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

func (c *consoleFormatter) Format(w io.Writer) (func(Token), error) {
	return func(token Token) {
		clr, ok := c.theme[token.Type]
		if !ok {
			clr, ok = c.theme[token.Type.SubCategory()]
			if !ok {
				clr, ok = c.theme[token.Type.Category()]
				if !ok {
					clr = "\033[0m"
				}
			}
		}
		fmt.Fprint(w, clr)
		fmt.Fprint(w, token.Value)
	}, nil
}
