package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CobolFree lexer.
var CobolFree = internal.Register(MustNewLexer(
	&Config{
		Name:      "COBOLFree",
		Aliases:   []string{"cobolfree"},
		Filenames: []string{"*.cbl", "*.CBL"},
	},
	Rules{
		"root": {},
	},
))
