package r

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var rslAnalyserRe = regexp.MustCompile(`(?i)scheme\s*.*?=\s*class\s*type`)

// RSL lexer. RSL <http://en.wikipedia.org/wiki/RAISE> is the formal
// specification language used in RAISE (Rigorous Approach to Industrial
// Software Engineering) method.
var RSL = internal.Register(MustNewLexer(
	&Config{
		Name:      "RSL",
		Aliases:   []string{"rsl"},
		Filenames: []string{"*.rsl"},
		MimeTypes: []string{"text/rsl"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// Check for the most common text in the beginning of a RSL file.
	if rslAnalyserRe.MatchString(text) {
		return 1.0
	}

	return 0
}))
