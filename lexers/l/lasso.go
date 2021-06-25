package l

import (
	"regexp"
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	lassoAnalyserDelimiterRe = regexp.MustCompile(`(?i)<\?lasso`)
	lassoAnalyserLocalRe     = regexp.MustCompile(`(?i)local\(`)
)

// Lasso lexer.
var Lasso = internal.Register(MustNewLexer(
	&Config{
		Name:    "Lasso",
		Aliases: []string{"lasso", "lassoscript"},
		Filenames: []string{
			"*.lasso",
			"*.lasso[89]",
		},
		AliasFilenames: []string{
			"*.incl",
			"*.inc",
			"*.las",
		},
		MimeTypes: []string{"text/x-lasso"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	var result float32

	if strings.Contains(text, "bin/lasso9") {
		result += 0.8
	}

	if lassoAnalyserDelimiterRe.MatchString(text) {
		result += 0.4
	}

	if lassoAnalyserLocalRe.MatchString(text) {
		result += 0.4
	}

	return result
}))
