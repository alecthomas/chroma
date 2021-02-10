package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Blazor lexer.
var Blazor = internal.Register(MustNewLexer(
	&Config{
		Name:      "Blazor",
		Aliases:   []string{"blazor"},
		Filenames: []string{"*.razor"},
	},
	Rules{
		"root": {},
	},
))
