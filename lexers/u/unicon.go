package u

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Unicon lexer.
var Unicon = internal.Register(MustNewLexer(
	&Config{
		Name:      "Unicon",
		Aliases:   []string{"unicon"},
		Filenames: []string{"*.icn"},
		MimeTypes: []string{"text/unicon"},
	},
	Rules{
		"root": {},
	},
))
