package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Chapel lexer.
var Chapel = internal.Register(MustNewLexer(
	&Config{
		Name:      "Chapel",
		Aliases:   []string{"chapel", "chpl"},
		Filenames: []string{"*.chpl"},
	},
	Rules{
		"root": {},
	},
))
