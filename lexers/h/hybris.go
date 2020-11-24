package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Hybris lexer.
var Hybris = internal.Register(MustNewLexer(
	&Config{
		Name:      "Hybris",
		Aliases:   []string{"hybris", "hy"},
		Filenames: []string{"*.hy", "*.hyb"},
		MimeTypes: []string{"text/x-hybris", "application/x-hybris"},
	},
	Rules{
		"root": {},
	},
))
