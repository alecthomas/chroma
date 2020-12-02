package i

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var inform6AnalyserRe = regexp.MustCompile(`(?i)\borigsource\b`)

// Inform6 lexer.
var Inform6 = internal.Register(MustNewLexer(
	&Config{
		Name:      "Inform 6",
		Aliases:   []string{"inform6", "i6"},
		Filenames: []string{"*.inf"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// We try to find a keyword which seem relatively common, unfortunately
	// there is a decent overlap with Smalltalk keywords otherwise here.
	if inform6AnalyserRe.MatchString(text) {
		return 0.05
	}

	return 0
}))
