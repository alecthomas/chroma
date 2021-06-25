package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pan lexer.
var Pan = internal.Register(MustNewLexer(
	&Config{
		Name:      "Pan",
		Aliases:   []string{"pan"},
		Filenames: []string{"*.pan"},
	},
	Rules{
		"root": {},
	},
))
