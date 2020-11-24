package p

import (
	"strings"

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
).SetAnalyser(func(text string) float32 {
	// This is basically C. There is a keyword which doesn't exist in C
	// though and is nearly unique to this language.
	if strings.Contains(text, "tagof") {
		return 0.01
	}

	return 0
}))
