package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Processing lexer.
var Processing = internal.Register(MustNewLexer(
	&Config{
		Name:      "Processing",
		Aliases:   []string{"processing"},
		Filenames: []string{"*.pde"},
	},
	Rules{
		"root": {},
	},
))
