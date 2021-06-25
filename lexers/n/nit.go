package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nit lexer.
var Nit = internal.Register(MustNewLexer(
	&Config{
		Name:      "Nit",
		Aliases:   []string{"nit"},
		Filenames: []string{"*.nit"},
	},
	Rules{
		"root": {},
	},
))
