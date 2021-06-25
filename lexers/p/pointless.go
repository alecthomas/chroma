package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pointless lexer.
var Pointless = internal.Register(MustNewLexer(
	&Config{
		Name:      "Pointless",
		Aliases:   []string{"pointless"},
		Filenames: []string{"*.ptls"},
	},
	Rules{
		"root": {},
	},
))
