package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ADL lexer.
var Adl = internal.Register(MustNewLexer(
	&Config{
		Name:      "ADL",
		Aliases:   []string{"adl"},
		Filenames: []string{"*.adl", "*.adls", "*.adlf", "*.adlx"},
	},
	Rules{
		"root": {},
	},
))
