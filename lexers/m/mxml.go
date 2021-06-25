package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MXML lexer.
var MXML = internal.Register(MustNewLexer(
	&Config{
		Name:      "MXML",
		Aliases:   []string{"mxml"},
		Filenames: []string{"*.mxml"},
		MimeTypes: []string{"text/xml", "application/xml"},
	},
	Rules{
		"root": {},
	},
))
