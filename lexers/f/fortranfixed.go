package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// FortranFixed lexer.
var FortranFixed = internal.Register(MustNewLexer(
	&Config{
		Name:      "FortranFixed",
		Aliases:   []string{"fortranfixed"},
		Filenames: []string{"*.f", "*.F"},
	},
	Rules{
		"root": {},
	},
))
