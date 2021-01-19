package l

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var logosAnalyserKeywordsRe = regexp.MustCompile(`%(?:hook|ctor|init|c\()`)

// Logos lexer.
var Logos = internal.Register(MustNewLexer(
	&Config{
		Name:      "Logos",
		Aliases:   []string{"logos"},
		Filenames: []string{"*.x", "*.xi", "*.xm", "*.xmi"},
		MimeTypes: []string{"text/x-logos"},
		Priority:  0.25,
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if logosAnalyserKeywordsRe.MatchString(text) {
		return 1.0
	}

	return 0
}))
