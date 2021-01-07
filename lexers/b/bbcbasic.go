package b

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// BBC Basic lexer.
var BbcBasic = internal.Register(MustNewLexer(
	&Config{
		Name:      "BBC Basic",
		Aliases:   []string{"bbcbasic"},
		Filenames: []string{"*.bbc"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if strings.HasPrefix(text, "10REM >") || strings.HasPrefix(text, "REM >") {
		return 0.9
	}

	return 0
}))
