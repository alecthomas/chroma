package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Execline lexer.
var Execline = internal.Register(MustNewLexer(
	&Config{
		Name:      "execline",
		Aliases:   []string{"execline"},
		Filenames: []string{"*.exec"},
	},
	Rules{
		"root": {},
	},
))
