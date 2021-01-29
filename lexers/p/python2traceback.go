package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Python2Traceback lexer.
var Python2Traceback = internal.Register(MustNewLexer(
	&Config{
		Name:      "Python 2.x Traceback",
		Aliases:   []string{"py2tb"},
		Filenames: []string{"*.py2tb"},
		MimeTypes: []string{"text/x-python2-traceback"},
	},
	Rules{
		"root": {},
	},
))
