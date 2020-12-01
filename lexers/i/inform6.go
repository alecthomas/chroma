package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Inform6 lexer.
var Inform6 = internal.Register(MustNewLexer(
	&Config{
		Name:      "Inform 6",
		Aliases:   []string{"inform6", "i6"},
		Filenames: []string{"*.inf"},
	},
	Rules{
		"root": {},
	},
))
