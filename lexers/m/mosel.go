package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Mosel lexer.
var Mosel = internal.Register(MustNewLexer(
	&Config{
		Name:      "Mosel",
		Aliases:   []string{"model"},
		Filenames: []string{"*.mos"},
	},
	Rules{
		"root": {},
	},
))
