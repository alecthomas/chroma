package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Evoque lexer.
var Evoque = internal.Register(MustNewLexer(
	&Config{
		Name:      "Evoque",
		Aliases:   []string{"evoque"},
		Filenames: []string{"*.evoque"},
		MimeTypes: []string{"application/x-evoque"},
	},
	Rules{
		"root": {},
	},
))
