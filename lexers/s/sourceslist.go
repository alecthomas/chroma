package s

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var sourcesListAnalyserRe = regexp.MustCompile(`(?m)^\s*(deb|deb-src) `)

// SourcesList lexer. Lexer that highlights debian sources.list files.
var SourcesList = internal.Register(MustNewLexer(
	&Config{
		Name:      "Debian Sourcelist",
		Aliases:   []string{"sourceslist", "sources.list", "debsources"},
		Filenames: []string{"sources.list"},
		MimeTypes: []string{"application/x-debian-sourceslist"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if sourcesListAnalyserRe.MatchString(text) {
		return 1.0
	}

	return 0
}))
