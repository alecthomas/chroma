package o

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var objectiveJAnalyserImportRe = regexp.MustCompile(`(?m)^\s*@import\s+[<"]`)

// Objective-J lexer.
var ObjectiveJ = internal.Register(MustNewLexer(
	&Config{
		Name:      "Objective-J",
		Aliases:   []string{"objective-j", "objectivej", "obj-j", "objj"},
		Filenames: []string{"*.j"},
		MimeTypes: []string{"text/x-objective-j"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// special directive found in most Objective-J files.
	if objectiveJAnalyserImportRe.MatchString(text) {
		return 1.0
	}

	return 0
}))
