package chroma

import (
	"fmt"
	"io"
	"iter"
)

// A Formatter for Chroma lexers.
type Formatter interface {
	// Format writes the tokens from iterator to w using style.
	//
	// If the iterator panics, the Formatter should recover and return the
	// panic as an error.
	Format(w io.Writer, style *Style, iterator iter.Seq[Token]) error
}

// A FormatterFunc is a Formatter implemented as a function.
//
// Guards against iterator panics.
type FormatterFunc func(w io.Writer, style *Style, iterator iter.Seq[Token]) error

func (f FormatterFunc) Format(w io.Writer, s *Style, it iter.Seq[Token]) (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			if e, ok := perr.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", perr)
			}
		}
	}()
	return f(w, s, it)
}
