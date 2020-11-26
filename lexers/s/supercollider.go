package s

import (
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
))
