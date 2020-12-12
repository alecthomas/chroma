package g

import (
	"math"
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	gapAnalyserDeclarationRe = regexp.MustCompile(
		`(InstallTrueMethod|Declare(Attribute|Category|Filter|Operation|GlobalFunction|Synonym|SynonymAttr|Property))`)
	gapAnalyserImplementationRe = regexp.MustCompile(
		`(DeclareRepresentation|Install(GlobalFunction|Method|ImmediateMethod|OtherMethod)|New(Family|Type)|Objectify)`)
)

// GAP lexer.
var Gap = internal.Register(MustNewLexer(
	&Config{
		Name:      "GAP",
		Aliases:   []string{"gap"},
		Filenames: []string{"*.g", "*.gd", "*.gi", "*.gap"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	var result float64

	if gapAnalyserDeclarationRe.MatchString(text) {
		result += 0.7
	}

	if gapAnalyserImplementationRe.MatchString(text) {
		result += 0.7
	}

	return float32(math.Min(result, float64(1.0)))
}))
