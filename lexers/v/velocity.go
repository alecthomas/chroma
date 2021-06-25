package v

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	velocityAnalzserMacroRe     = regexp.MustCompile(`(?s)#\{?macro\}?\(.*?\).*?#\{?end\}?`)
	velocityAnalzserIfRe        = regexp.MustCompile(`(?s)#\{?if\}?\(.+?\).*?#\{?end\}?`)
	velocityAnalzserForeachRe   = regexp.MustCompile(`(?s)#\{?foreach\}?\(.+?\).*?#\{?end\}?`)
	velocityAnalzserReferenceRe = regexp.MustCompile(`\$!?\{?[a-zA-Z_]\w*(\([^)]*\))?(\.\w+(\([^)]*\))?)*\}?`)
)

// Velocity lexer.
var Velocity = internal.Register(MustNewLexer(
	&Config{
		Name:      "Velocity",
		Aliases:   []string{"velocity"},
		Filenames: []string{"*.vm", "*.fhtml"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	var result float64

	if velocityAnalzserMacroRe.MatchString(text) {
		result += 0.25
	}

	if velocityAnalzserIfRe.MatchString(text) {
		result += 0.15
	}

	if velocityAnalzserForeachRe.MatchString(text) {
		result += 0.15
	}

	if velocityAnalzserReferenceRe.MatchString(text) {
		result += 0.01
	}

	return float32(result)
}))
