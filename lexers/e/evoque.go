package e

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Evoque lexer.
var Evoque = internal.Register(MustNewLexer(
	&Config{
		Name:      "Evoque",
		Aliases:   []string{"evoque"},
		Filenames: []string{"*.evoque"},
		MimeTypes: []string{"application/x-evoque"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// Evoque templates use $evoque, which is unique.
	if strings.Contains(text, "$evoque") {
		return 1.0
	}

	return 0
}))
