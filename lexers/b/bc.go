package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Bc lexer.
var Bc = internal.Register(MustNewLexer(
	&Config{
		Name:      "BC",
		Aliases:   []string{"bc"},
		Filenames: []string{"*.bc"},
	},
	Rules{
		"root": {},
	},
))
