package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Xtlang lexer.
var Xtlang = internal.Register(MustNewLexer(
	&Config{
		Name:      "xtlang",
		Aliases:   []string{"extempore"},
		Filenames: []string{"*.xtm"},
	},
	Rules{
		"root": {},
	},
))
