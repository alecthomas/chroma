package j

import (
	"regexp"
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var jclAnalyserJobHeaderRe = regexp.MustCompile(`(?i)^//[a-z#$@][a-z0-9#$@]{0,7}\s+job(\s+.*)?$`)

// Jcl lexer.
var Jcl = internal.Register(MustNewLexer(
	&Config{
		Name:      "JCL",
		Aliases:   []string{"jcl"},
		Filenames: []string{"*.jcl"},
		MimeTypes: []string{"text/x-jcl"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// Recognize JCL job by header.
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return 0
	}

	if jclAnalyserJobHeaderRe.MatchString(lines[0]) {
		return 1.0
	}

	return 0
}))
