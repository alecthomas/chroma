package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TADS 3 lexer.
var Tads3 = internal.Register(MustNewLexer(
	&Config{
		Name:      "TADS 3",
		Aliases:   []string{"tads3"},
		Filenames: []string{"*.t"},
	},
	Rules{
		"root": {},
	},
))
