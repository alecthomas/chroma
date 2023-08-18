package lexers

import (
	"regexp"

	. "github.com/alecthomas/chroma/v2" // nolint
)

var (
	cAnalyserIncludeRe = regexp.MustCompile(`(?m)^\s*#include [<"]`)
	cAnalyserIfdefRe   = regexp.MustCompile(`(?m)^\s*#ifn?def `)
)

// C lexer.
var C = Register(MustNewXMLLexer(
	embedded,
	"embedded/c.xml",
).SetConfig(
	&Config{
		Name:      "C",
		Aliases:   []string{"c"},
		Filenames: []string{"*.c", "*.h", "*.idc", "*.x[bp]m"},
		MimeTypes: []string{"text/x-chdr", "text/x-csrc", "image/x-xbitmap", "image/x-xpixmap"},
		EnsureNL:  true,
		Priority:  0.1,
	},
).SetAnalyser(func(text string) float32 {
	if cAnalyserIncludeRe.MatchString(text) {
		return 0.1
	}

	if cAnalyserIfdefRe.MatchString(text) {
		return 0.1
	}

	return 0
}))
