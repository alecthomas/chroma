package c

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var ca65AnalyserCommentRe = regexp.MustCompile(`(?m)^\s*;`)

// Ca65 lexer.
var Ca65 = internal.Register(MustNewLexer(
	&Config{
		Name:      "ca65 assembler",
		Aliases:   []string{"ca65"},
		Filenames: []string{"*.s"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// comments in GAS start with "#".
	if ca65AnalyserCommentRe.MatchString(text) {
		return 0.9
	}

	return 0
}))
