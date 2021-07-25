package j

import (
	"math"
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	jasminAnalyserClassRe       = regexp.MustCompile(`(?m)^\s*\.class\s`)
	jasminAnalyserInstructionRe = regexp.MustCompile(`(?m)^\s*[a-z]+_[a-z]+\b`)
	jasminAnalyserKeywordsRe    = regexp.MustCompile(
		`(?m)^\s*\.(attribute|bytecode|debug|deprecated|enclosing|inner|interface|limit|set|signature|stack)\b`)
)

// Jasmin lexer.
var Jasmin = internal.Register(MustNewLexer(
	&Config{
		Name:      "Jasmin",
		Aliases:   []string{"jasmin", "jasminxt"},
		Filenames: []string{"*.j"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	var result float64

	if jasminAnalyserClassRe.MatchString(text) {
		result += 0.5

		if jasminAnalyserInstructionRe.MatchString(text) {
			result += 0.3
		}
	}

	if jasminAnalyserKeywordsRe.MatchString(text) {
		result += 0.6
	}

	return float32(math.Min(result, float64(1.0)))
}))
