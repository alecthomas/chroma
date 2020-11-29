package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Velocity lexer.
var Velocity = internal.Register(MustNewLexer(
	&Config{
		Name:      "Velocity",
		Aliases:   []string{"velocity"},
		Filenames: []string{"*.vm", "*.fhtml"},
	},
	Rules{
		"root": {},
	},
))
