package n

import (
	"github.com/alecthomas/chroma"
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/c"
	"github.com/alecthomas/chroma/lexers/internal"
)

// NesC lexer.
var NesC = internal.Register(MustNewLexer(
	&Config{
		Name:      "nesC",
		Aliases:   []string{"nesc"},
		Filenames: []string{"*.nc"},
		MimeTypes: []string{"text/x-nescsrc"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if analyser, ok := c.C.(chroma.Analyser); ok {
		return analyser.AnalyseText(text)
	}

	return 0
}))
