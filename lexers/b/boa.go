package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Boa lexer.
var Boa = internal.Register(MustNewLexer(
	&Config{
		Name:      "Boa",
		Aliases:   []string{"boa"},
		Filenames: []string{"*.boa"},
	},
	Rules{
		"root": {},
	},
))
