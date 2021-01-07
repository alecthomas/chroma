package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Boogie lexer.
var Boogie = internal.Register(MustNewLexer(
	&Config{
		Name:      "Boogie",
		Aliases:   []string{"boogie"},
		Filenames: []string{"*.bpl"},
	},
	Rules{
		"root": {},
	},
))
