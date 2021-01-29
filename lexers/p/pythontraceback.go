package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PythonTraceback lexer.
var PythonTraceback = internal.Register(MustNewLexer(
	&Config{
		Name:      "Python Traceback",
		Aliases:   []string{"pytb", "py3tb"},
		Filenames: []string{"*.pytb", "*.py3tb"},
		MimeTypes: []string{"text/x-python-traceback", "text/x-python3-traceback"},
	},
	Rules{
		"root": {},
	},
))
