package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Golo lexer.
var Golo = internal.Register(MustNewLexer(
	&Config{
		Name:      "Golo",
		Aliases:   []string{"golo"},
		Filenames: []string{"*.golo"},
	},
	Rules{
		"root": {},
	},
))
