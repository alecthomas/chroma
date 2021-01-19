package j

import (
	"strings"

	"github.com/alecthomas/chroma"
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/xml"
)

// Jsp lexer.
var Jsp = internal.Register(MustNewLexer(
	&Config{
		Name:      "Java Server Page",
		Aliases:   []string{"jsp"},
		Filenames: []string{"*.jsp"},
		MimeTypes: []string{"application/x-jsp"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	var result float32

	if analyser, ok := Java.(chroma.Analyser); ok {
		result = analyser.AnalyseText(text) - 0.01
	}

	if xml.MatchString(text) {
		result += 0.4
	}

	if strings.Contains(text, "<%") && strings.Contains(text, "%>") {
		result += 0.1
	}

	return result
}))
