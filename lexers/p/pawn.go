package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pawn lexer.
var Pawn = internal.Register(MustNewLexer(
	&Config{
		Name:      "Pawn",
		Aliases:   []string{"pawn"},
		Filenames: []string{"*.p", "*.pwn", "*.inc"},
		MimeTypes: []string{"text/x-pawn"},
	},
	Rules{
		"root": {},
	},
))
