package n

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Notmuch lexer.
var Notmuch = internal.Register(MustNewLexer(
	&Config{
		Name:    "Notmuch",
		Aliases: []string{"notmuch"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if strings.HasPrefix(text, "\fmessage{") {
		return 1.0
	}

	return 0
}))
