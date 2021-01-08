package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cadl lexer.
var Cadl = internal.Register(MustNewLexer(
	&Config{
		Name:      "cADL",
		Aliases:   []string{"cadl"},
		Filenames: []string{"*.cadl"},
	},
	Rules{
		"root": {},
	},
))
