package formatters

import (
	"io"

	"github.com/alecthomas/chroma"
)

// NoOp formatter.
var NoOp = Register("noop", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style) (func(*chroma.Token), error) {
	return func(t *chroma.Token) { io.WriteString(w, t.Value) }, nil
}))

// Fallback formatter.
var Fallback = NoOp

// Registry of Formatters.
var Registry = map[string]chroma.Formatter{}

// Get formatter by name.
//
// If the given formatter is not found, the Fallback formatter will be returned.
func Get(name string) chroma.Formatter {
	if f, ok := Registry[name]; ok {
		return f
	}
	return Fallback
}

// Register a named formatter.
func Register(name string, formatter chroma.Formatter) chroma.Formatter {
	Registry[name] = formatter
	return formatter
}
