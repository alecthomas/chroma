package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/shebang"
)

// Execline lexer.
var Execline = internal.Register(MustNewLexer(
	&Config{
		Name:      "execline",
		Aliases:   []string{"execline"},
		Filenames: []string{"*.exec"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if matched, _ := shebang.MatchString(text, "execlineb"); matched {
		return 1.0
	}

	return 0
}))
