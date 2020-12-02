package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Jasmin lexer.
var Jasmin = internal.Register(MustNewLexer(
	&Config{
		Name:      "Jasmin",
		Aliases:   []string{"jasmin", "jasminxt"},
		Filenames: []string{"*.j"},
	},
	Rules{
		"root": {},
	},
))
