package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Liquid lexer.
var Liquid = internal.Register(MustNewLexer(
	&Config{
		Name:      "liquid",
		Aliases:   []string{"liquid"},
		Filenames: []string{"*.liquid"},
	},
	Rules{
		"root": {},
	},
))
