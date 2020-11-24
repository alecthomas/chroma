package h

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var hybrisAnalyserRe = regexp.MustCompile(`\b(?:public|private)\s+method\b`)

// Hybris lexer.
var Hybris = internal.Register(MustNewLexer(
	&Config{
		Name:      "Hybris",
		Aliases:   []string{"hybris", "hy"},
		Filenames: []string{"*.hy", "*.hyb"},
		MimeTypes: []string{"text/x-hybris", "application/x-hybris"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// public method and private method don't seem to be quite common
	// elsewhere.
	if hybrisAnalyserRe.MatchString(text) {
		return 0.01
	}

	return 0
}))
