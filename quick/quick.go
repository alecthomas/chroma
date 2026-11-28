// Package quick provides simple, no-configuration source code highlighting.
package quick

import (
	"io"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/formatters"
	"github.com/alecthomas/chroma/v3/lexers"
	"github.com/alecthomas/chroma/v3/styles"
)

// Highlight some text.
//
// Lexer, formatter and style may be empty, in which case a best-effort is made.
func Highlight(w io.Writer, source, lexer, formatter, style string) error {
	// Determine lexer.
	l := lexers.Get(lexer)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	f := formatters.Get(formatter)
	s := styles.Get(style)
	it, err := l.Tokenise(source)
	if err != nil {
		return err
	}
	return f.Format(w, s, it)
}
