package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Jsgf lexer.
var Jsgf = internal.Register(MustNewLexer(
	&Config{
		Name:      "JSGF",
		Aliases:   []string{"jsgf"},
		Filenames: []string{"*.jsgf"},
		MimeTypes: []string{"application/jsgf", "application/x-jsgf", "text/jsgf"},
	},
	Rules{
		"root": {},
	},
))
