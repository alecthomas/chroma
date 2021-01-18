package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Inform7 lexer.
var Inform7 = internal.Register(MustNewLexer(
	&Config{
		Name:      "Inform 7",
		Aliases:   []string{"inform7", "i7"},
		Filenames: []string{"*.ni", "*.i7x"},
	},
	Rules{
		"root": {},
	},
))
