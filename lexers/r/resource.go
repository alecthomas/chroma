package r

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Resource lexer. Lexer for ICU Resource bundles
// <http://userguide.icu-project.org/locale/resources>
var Resource = internal.Register(MustNewLexer(
	&Config{
		Name:    "ResourceBundle",
		Aliases: []string{"resource", "resourcebundle"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if strings.HasPrefix(text, "root:table") {
		return 1.0
	}

	return 0
}))
