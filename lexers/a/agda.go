package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Agda lexer.
var Agda = internal.Register(MustNewLexer(
	&Config{
		Name:      "Agda",
		Aliases:   []string{"agda"},
		Filenames: []string{"*.agda"},
		MimeTypes: []string{"text/x-agda"},
	},
	Rules{
		"root": {},
	},
))
