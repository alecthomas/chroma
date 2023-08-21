package lexers

import (
	"regexp"

	. "github.com/alecthomas/chroma/v2" // nolint
)

var actionscript3AnalyserRe = regexp.MustCompile(`\w+\s*:\s*\w`)

// Actionscript 3 lexer.
var Actionscript3 = Register(MustNewXMLLexer(
	embedded,
	"embedded/actionscript_3.xml",
).SetConfig(
	&Config{
		Name:      "ActionScript 3",
		Aliases:   []string{"as3", "actionscript3"},
		Filenames: []string{"*.as"},
		MimeTypes: []string{"application/x-actionscript3", "text/x-actionscript3", "text/actionscript3"},
		DotAll:    true,
	},
).SetAnalyser(func(text string) float32 {
	if actionscript3AnalyserRe.MatchString(text) {
		return 0.3
	}

	return 0
}))
