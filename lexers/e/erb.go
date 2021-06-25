package e

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Erb lexer.
var Erb = internal.Register(MustNewLexer(
	&Config{
		Name:      "ERB",
		Aliases:   []string{"erb"},
		MimeTypes: []string{"application/x-ruby-templating"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, "<%") && strings.Contains(text, "%>") {
		return 0.4
	}

	return 0
}))
