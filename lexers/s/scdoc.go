package s

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// scdoc lexer.
var Scdoc = internal.Register(MustNewLexer(
	&Config{
		Name:      "scdoc",
		Aliases:   []string{"scdoc", "scd"},
		Filenames: []string{"*.scd", "*.scdoc"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// This is very similar to markdown, save for the escape characters
	// needed for * and _.
	var result float32

	if strings.Contains(text, `\*`) {
		result += 0.01
	}

	if strings.Contains(text, `\_`) {
		result += 0.01
	}

	return result
}))
