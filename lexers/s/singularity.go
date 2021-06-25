package s

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	singularityAnalyserHeaderRe  = regexp.MustCompile(`(?i)\b(?:osversion|includecmd|mirrorurl)\b`)
	singularityAnalyserSectionRe = regexp.MustCompile(`%(?:pre|post|setup|environment|help|labels|test|runscript|files|startscript)\b`)
)

// Singularity lexer.
var Singularity = internal.Register(MustNewLexer(
	&Config{
		Name:      "Singularity",
		Aliases:   []string{"singularity"},
		Filenames: []string{"*.def", "Singularity"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// This is a quite simple script file, but there are a few keywords
	// which seem unique to this language.
	var result float32

	if singularityAnalyserHeaderRe.MatchString(text) {
		result += 0.5
	}

	if singularityAnalyserSectionRe.MatchString(text) {
		result += 0.49
	}

	return result
}))
