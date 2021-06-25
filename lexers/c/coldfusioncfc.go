package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ColdfusionCfc lexer.
var ColdfusionCfc = internal.Register(MustNewLexer(
	&Config{
		Name:      "Coldfusion CFC",
		Aliases:   []string{"cfc"},
		Filenames: []string{"*.cfc"},
	},
	Rules{
		"root": {},
	},
))
