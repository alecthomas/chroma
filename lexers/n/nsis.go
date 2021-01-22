package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// NSIS lexer.
var NSIS = internal.Register(MustNewLexer(
	&Config{
		Name:      "NSIS",
		Aliases:   []string{"nsis", "nsi", "nsh"},
		Filenames: []string{"*.nsi", "*.nsh"},
		MimeTypes: []string{"text/x-nsis"},
	},
	Rules{
		"root": {},
	},
))
