package j

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	jagsAnalyserModelRe = regexp.MustCompile(`(?m)^\s*model\s*\{`)
	jagsAnalyserDataRe  = regexp.MustCompile(`(?m)^\s*data\s*\{`)
	jagsAnalyserVarRe   = regexp.MustCompile(`(?m)^\s*var`)
)

// JAGS lexer.
var Jags = internal.Register(MustNewLexer(
	&Config{
		Name:      "JAGS",
		Aliases:   []string{"jags"},
		Filenames: []string{"*.jag", "*.bug"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if jagsAnalyserModelRe.MatchString(text) {
		if jagsAnalyserDataRe.MatchString(text) {
			return 0.9
		}

		if jagsAnalyserVarRe.MatchString(text) {
			return 0.9
		}

		return 0.3
	}

	return 0
}))
