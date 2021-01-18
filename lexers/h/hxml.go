package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Hxml lexer.
var Hxml = internal.Register(MustNewLexer(
	&Config{
		Name:      "Hxml",
		Aliases:   []string{"haxeml", "hxml"},
		Filenames: []string{"*.hxml"},
	},
	Rules{
		"root": {},
	},
))
