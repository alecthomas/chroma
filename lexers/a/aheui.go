package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Aheui lexer.
var Aheui = internal.Register(MustNewLexer(
	&Config{
		Name:      "Aheui",
		Aliases:   []string{"aheui"},
		Filenames: []string{"*.aheui"},
	},
	Rules{
		"root": {},
	},
))
