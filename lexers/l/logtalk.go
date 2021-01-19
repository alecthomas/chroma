package l

import (
	"regexp"
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var logtalkAnalyserSyntaxRe = regexp.MustCompile(`(?m)^:-\s[a-z]`)

// Logtalk lexer.
var Logtalk = internal.Register(MustNewLexer(
	&Config{
		Name:      "Logtalk",
		Aliases:   []string{"logtalk"},
		Filenames: []string{"*.lgt", "*.logtalk"},
		MimeTypes: []string{"text/x-logtalk"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, ":- object(") ||
		strings.Contains(text, ":- protocol(") ||
		strings.Contains(text, ":- category(") {
		return 1.0
	}

	if logtalkAnalyserSyntaxRe.MatchString(text) {
		return 0.9
	}

	return 0
}))
