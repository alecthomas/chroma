package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Marko lexer.
var Marko = internal.Register(MustNewLexer(
	&Config{
		Name:      "Marko",
		Aliases:   []string{"marko"},
		Filenames: []string{"*.marko"},
		MimeTypes: []string{"text/x-marko"},
	},
	Rules{
		"root": {},
	},
))
