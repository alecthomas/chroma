package chroma

import (
	"io"
)

// A Formatter for Chroma lexers.
type Formatter interface {
	// Format returns a formatting function for tokens.
	Format(w io.Writer, style *Style, iterator Iterator) error
}

// A FormatterFunc is a Formatter implemented as a function.
type FormatterFunc func(w io.Writer, style *Style, iterator Iterator) error

func (f FormatterFunc) Format(w io.Writer, s *Style, it Iterator) error { return f(w, s, it) }
