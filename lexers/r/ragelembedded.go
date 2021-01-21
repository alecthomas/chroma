package r

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RagelEmbedded lexer. A lexer for Ragel embedded in a host language file.
var RagelEmbedded = internal.Register(MustNewLexer(
	&Config{
		Name:      "Embedded Ragel",
		Aliases:   []string{"ragel-em"},
		Filenames: []string{"*.rl"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, "@LANG: indep") {
		return 1.0
	}

	return 0
}))
