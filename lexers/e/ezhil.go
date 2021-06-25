package e

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var ezhilAnalyserRe = regexp.MustCompile(`[u0b80-u0bff]`)

// Ezhil lexer.
var Ezhil = internal.Register(MustNewLexer(
	&Config{
		Name:      "Ezhil",
		Aliases:   []string{"ezhil"},
		Filenames: []string{"*.n"},
		MimeTypes: []string{"text/x-ezhil"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// this language uses Tamil-script. We'll assume that if there's a
	// decent amount of Tamil-characters, it's this language. This assumption
	// is obviously horribly off if someone uses string literals in tamil
	// in another language.
	if len(ezhilAnalyserRe.FindAllString(text, -1)) > 10 {
		return 0.25
	}

	return 0
}))
