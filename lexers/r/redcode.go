package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Redcode lexer.
var Redcode = internal.Register(MustNewLexer(
	&Config{
		Name:      "Redcode",
		Aliases:   []string{"redcode"},
		Filenames: []string{"*.cw"},
	},
	Rules{
		"root": {},
	},
))
