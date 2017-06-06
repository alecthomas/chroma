package formatters

import (
	"io"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/styles"
)

// A Formatter for Chroma lexers.
type Formatter interface {
	// Format returns a formatting function for tokens.
	Format(w io.Writer, style *styles.Style) (func(*chroma.Token), error)
}

type FormatterFunc func(io.Writer, *styles.Style) (func(*chroma.Token), error)

func (f FormatterFunc) Format(w io.Writer, s *styles.Style) (func(*chroma.Token), error) {
	return f(w, s)
}

var noop = Register("noop", FormatterFunc(func(w io.Writer, s *styles.Style) (func(*chroma.Token), error) {
	return func(t *chroma.Token) { io.WriteString(w, t.Value) }, nil
}))

// Fallback formatter.
var Fallback = noop

// Registry of Formatters.
var Registry = map[string]Formatter{}

// Get formatter by name.
//
// If the given formatter is not found, the Fallback formatter will be returned.
func Get(name string) Formatter {
	if f, ok := Registry[name]; ok {
		return f
	}
	return Fallback
}

// Register a named formatter.
func Register(name string, formatter Formatter) Formatter {
	Registry[name] = formatter
	return formatter
}
