package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Clay lexer.
var Clay = internal.Register(MustNewLexer(
	&Config{
		Name:      "Clay",
		Aliases:   []string{"clay"},
		Filenames: []string{"*.clay"},
		MimeTypes: []string{"text/x-clay"},
	},
	Rules{
		"root": {},
	},
))
