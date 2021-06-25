package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// FoxPro lexer.
var FoxPro = internal.Register(MustNewLexer(
	&Config{
		Name:      "FoxPro",
		Aliases:   []string{"foxpro", "vfp", "clipper", "xbase"},
		Filenames: []string{"*.PRG", "*.prg"},
	},
	Rules{
		"root": {},
	},
))
