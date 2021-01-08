package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cpsa lexer.
var Cpsa = internal.Register(MustNewLexer(
	&Config{
		Name:      "CPSA",
		Aliases:   []string{"cpsa"},
		Filenames: []string{"*.cpsa"},
	},
	Rules{
		"root": {},
	},
))
