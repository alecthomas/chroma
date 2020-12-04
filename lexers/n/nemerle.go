package n

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nemerle lexer.
var Nemerle = internal.Register(MustNewLexer(
	&Config{
		Name:      "Nemerle",
		Aliases:   []string{"nemerle"},
		Filenames: []string{"*.n"},
		// inferred
		MimeTypes: []string{"text/x-nemerle"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// Nemerle is quite similar to Python, but @if is relatively uncommon
	// elsewhere.
	if strings.Contains(text, "@if") {
		return 0.1
	}

	return 0
}))
