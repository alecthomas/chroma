package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Hspec lexer.
var Hspec = internal.Register(MustNewLexer(
	&Config{
		Name:    "Hspec",
		Aliases: []string{"hspec"},
	},
	Rules{
		"root": {},
	},
))
