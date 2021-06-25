package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PyPyLog lexer.
var PyPyLog = internal.Register(MustNewLexer(
	&Config{
		Name:      "PyPy Log",
		Aliases:   []string{"pypylog", "pypy"},
		Filenames: []string{"*.pypylog"},
		MimeTypes: []string{"application/x-pypylog"},
	},
	Rules{
		"root": {},
	},
))
