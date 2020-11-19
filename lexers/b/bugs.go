package b

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var bugsAnalyzerRe = regexp.MustCompile(`(?m)^\s*model\s*{`)

// BUGS lexer.
var Bugs = internal.Register(MustNewLexer(
	&Config{
		Name:      "BUGS",
		Aliases:   []string{"bugs", "winbugs", "openbugs"},
		Filenames: []string{"*.bug"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if bugsAnalyzerRe.MatchString(text) {
		return 0.7
	}

	return 0
}))
