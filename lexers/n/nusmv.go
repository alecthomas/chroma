package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// NuSMV lexer.
var NuSMV = internal.Register(MustNewLexer(
	&Config{
		Name:      "NuSMV",
		Aliases:   []string{"nusmv"},
		Filenames: []string{"*.smv"},
	},
	Rules{
		"root": {},
	},
))
