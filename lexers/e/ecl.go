package e

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ECL lexer.
var Ecl = internal.Register(MustNewLexer(
	&Config{
		Name:      "ECL",
		Aliases:   []string{"ecl"},
		Filenames: []string{"*.ecl"},
		MimeTypes: []string{"application/x-ecl"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// This is very difficult to guess relative to other business languages.
	// -> in conjunction with BEGIN/END seems relatively rare though.

	var result float32 = 0

	if strings.Contains(text, "->") {
		result += 0.01
	}

	if strings.Contains(text, "BEGIN") {
		result += 0.01
	}

	if strings.Contains(text, "END") {
		result += 0.01
	}

	return result
}))
