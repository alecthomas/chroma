package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Arrow lexer.
var Arrow = internal.Register(MustNewLexer(
	&Config{
		Name:      "Arrow",
		Aliases:   []string{"arrow"},
		Filenames: []string{"*.arw"},
	},
	Rules{
		"root": {},
	},
))
