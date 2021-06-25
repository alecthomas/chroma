package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Charmci lexer.
var Charmci = internal.Register(MustNewLexer(
	&Config{
		Name:      "Charmci",
		Aliases:   []string{"charmci"},
		Filenames: []string{"*.ci"},
	},
	Rules{
		"root": {},
	},
))
