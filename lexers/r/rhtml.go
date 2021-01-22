package r

import (
	"github.com/alecthomas/chroma"
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/e"
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/doctype"
)

// RHTML lexer. Subclass of the ERB lexer that highlights the unlexed data
// with the html lexer.
var RHTML = internal.Register(MustNewLexer(
	&Config{
		Name:           "RHTML",
		Aliases:        []string{"rhtml", "html+erb", "html+ruby"},
		Filenames:      []string{"*.rhtml"},
		AliasFilenames: []string{"*.html", "*.htm", "*.xhtml"},
		MimeTypes:      []string{"text/html+ruby"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	analyser, ok := e.Erb.(chroma.Analyser)
	if !ok {
		return 0
	}

	result := analyser.AnalyseText(text) - 0.01

	if matched, _ := doctype.MatchString(text, "html"); matched {
		// one more than the XmlErbLexer returns
		result += 0.5
	}

	return result
}))
