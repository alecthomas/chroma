package chroma

import (
	"io"
	"iter"
)

// A Formatter for Chroma lexers.
type Formatter interface {
	// Format returns a formatting function for tokens.
	//
	// If the iterator panics, the Formatter should recover.
	Format(w io.Writer, style *Style, iterator iter.Seq[Token]) error
}

// A FormatterFunc is a Formatter implemented as a function.
//
// Guards against iterator panics.
type FormatterFunc func(w io.Writer, style *Style, iterator iter.Seq[Token]) error

func (f FormatterFunc) Format(w io.Writer, s *Style, it iter.Seq[Token]) (err error) { // nolint
	defer func() {
		if perr := recover(); perr != nil {
			err = perr.(error)
		}
	}()
	return f(w, s, it)
}

type recoveringFormatter struct {
	Formatter
}

func (r recoveringFormatter) Format(w io.Writer, s *Style, it iter.Seq[Token]) (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = perr.(error)
		}
	}()
	return r.Formatter.Format(w, s, it)
}

// RecoveringFormatter wraps a formatter with panic recovery.
func RecoveringFormatter(formatter Formatter) Formatter { return recoveringFormatter{formatter} }
