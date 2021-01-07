package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Augeas lexer.
var Augeas = internal.Register(MustNewLexer(
	&Config{
		Name:      "Augeas",
		Aliases:   []string{"augeas"},
		Filenames: []string{"*.aug"},
	},
	Rules{
		"root": {},
	},
))
