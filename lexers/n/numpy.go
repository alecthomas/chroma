package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// NumPy lexer.
var NumPy = internal.Register(MustNewLexer(
	&Config{
		Name:    "NumPy",
		Aliases: []string{"numpy"},
	},
	Rules{
		"root": {},
	},
))
