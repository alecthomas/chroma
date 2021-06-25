package c

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var cbmBasicV2AnalyserRe = regexp.MustCompile(`^\d+`)

// CBM BASIC V2 lexer.
var CbmBasicV2 = internal.Register(MustNewLexer(
	&Config{
		Name:      "CBM BASIC V2",
		Aliases:   []string{"cbmbas"},
		Filenames: []string{"*.bas"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// if it starts with a line number, it shouldn't be a "modern" Basic
	// like VB.net
	if cbmBasicV2AnalyserRe.MatchString(text) {
		return 0.2
	}

	return 0
}))
