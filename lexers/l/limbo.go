package l

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var limboAnalyzerRe = regexp.MustCompile(`(?m)^implement \w+;`)

// Limbo lexer.
var Limbo = internal.Register(MustNewLexer(
	&Config{
		Name:      "Limbo",
		Aliases:   []string{"limbo"},
		Filenames: []string{"*.b"},
		MimeTypes: []string{"text/limbo"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// Any limbo module implements something
	if limboAnalyzerRe.MatchString(text) {
		return 0.7
	}

	return 0
}))
