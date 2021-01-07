package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ampl lexer.
var Ampl = internal.Register(MustNewLexer(
	&Config{
		Name:      "Ampl",
		Aliases:   []string{"ampl"},
		Filenames: []string{"*.run"},
	},
	Rules{
		"root": {},
	},
))
