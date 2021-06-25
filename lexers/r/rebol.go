package r

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	rebolAnalyserHeaderRe              = regexp.MustCompile(`^\s*REBOL\s*\[`)
	rebolAnalyserHeaderPrecedingTextRe = regexp.MustCompile(`\s*REBOL\s*\[`)
)

// Rebol lexer.
var Rebol = internal.Register(MustNewLexer(
	&Config{
		Name:      "REBOL",
		Aliases:   []string{"rebol"},
		Filenames: []string{"*.r", "*.r3", "*.reb"},
		MimeTypes: []string{"text/x-rebol"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// Check if code contains REBOL header, then it's probably not R code
	if rebolAnalyserHeaderRe.MatchString(text) {
		return 1.0
	}

	if rebolAnalyserHeaderPrecedingTextRe.MatchString(text) {
		return 0.5
	}

	return 0
}))
