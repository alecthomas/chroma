package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Clean lexer.
var Clean = internal.Register(MustNewLexer(
	&Config{
		Name:      "Clean",
		Aliases:   []string{"clean"},
		Filenames: []string{"*.icl", "*.dcl"},
	},
	Rules{
		"root": {},
	},
))
