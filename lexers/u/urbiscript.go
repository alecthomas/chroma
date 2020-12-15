package u

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// UrbiScript lexer.
var UrbiScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "UrbiScript",
		Aliases:   []string{"urbiscript"},
		Filenames: []string{"*.u"},
		MimeTypes: []string{"application/x-urbiscript"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// This is fairly similar to C and others, but freezeif and
	// waituntil are unique keywords.
	var result float32

	if strings.Contains(text, "freezeif") {
		result += 0.05
	}

	if strings.Contains(text, "waituntil") {
		result += 0.05
	}

	return result
}))
