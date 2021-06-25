package s

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SuperCollider lexer.
var SuperCollider = internal.Register(MustNewLexer(
	&Config{
		Name:      "SuperCollider",
		Aliases:   []string{"sc", "supercollider"},
		Filenames: []string{"*.sc", "*.scd"},
		MimeTypes: []string{"application/supercollider", "text/supercollider"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// We're searching for a common function and a unique keyword here.
	if strings.Contains(text, "SinOsc") || strings.Contains(text, "thisFunctionDef") {
		return 0.1
	}

	return 0
}))
