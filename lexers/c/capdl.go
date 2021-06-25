package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Capdl lexer.
var Capdl = internal.Register(MustNewLexer(
	&Config{
		Name:      "CapDL",
		Aliases:   []string{"capdl"},
		Filenames: []string{"*.cdl"},
	},
	Rules{
		"root": {},
	},
))
