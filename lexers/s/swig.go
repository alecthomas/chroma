package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SWIG lexer.
var SWIG = internal.Register(MustNewLexer(
	&Config{
		Name:      "SWIG",
		Aliases:   []string{"swig"},
		Filenames: []string{"*.swg", "*.i"},
		MimeTypes: []string{"text/swig"},
		// Lower than C/C++ and Objective C/C++
		Priority: 0.04,
	},
	Rules{
		"root": {},
	},
))
