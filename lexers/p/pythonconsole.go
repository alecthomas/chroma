package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PythonConsole lexer.
var PythonConsole = internal.Register(MustNewLexer(
	&Config{
		Name:      "Python console session",
		Aliases:   []string{"pycon"},
		MimeTypes: []string{"text/x-python-doctest"},
	},
	Rules{
		"root": {},
	},
))
