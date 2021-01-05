package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Alloy lexer.
var Alloy = internal.Register(MustNewLexer(
	&Config{
		Name:      "Alloy",
		Aliases:   []string{"alloy"},
		Filenames: []string{"*.als"},
		MimeTypes: []string{"text/x-alloy"},
	},
	Rules{
		"root": {},
	},
))
