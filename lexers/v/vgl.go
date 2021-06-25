package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VGL lexer.
var VGL = internal.Register(MustNewLexer(
	&Config{
		Name:      "VGL",
		Aliases:   []string{"vgl"},
		Filenames: []string{"*.rpf"},
	},
	Rules{
		"root": {},
	},
))
