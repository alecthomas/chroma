package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// BlitzMax lexer.
var BlitzMax = internal.Register(MustNewLexer(
	&Config{
		Name:      "BlitzMax",
		Aliases:   []string{"blitzmax", "bmax"},
		Filenames: []string{"*.bmx"},
		MimeTypes: []string{"text/x-bmx"},
	},
	Rules{
		"root": {},
	},
))
