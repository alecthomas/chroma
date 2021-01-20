package s

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var stanAnalyserRe = regexp.MustCompile(`(?m)^\s*parameters\s*\{`)

// Stan lexer. Lexer for Stan models.
//
// The Stan modeling language is specified in the *Stan Modeling Language
// User's Guide and Reference Manual, v2.17.0*,
// pdf <https://github.com/stan-dev/stan/releases/download/v2.17.0/stan-reference-2.17.0.pdf>`.
var Stan = internal.Register(MustNewLexer(
	&Config{
		Name:      "Stan",
		Aliases:   []string{"stan"},
		Filenames: []string{"*.stan"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if stanAnalyserRe.MatchString(text) {
		return 1.0
	}

	return 0
}))
