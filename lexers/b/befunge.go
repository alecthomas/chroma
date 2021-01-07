package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Befunge lexer.
var Befunge = internal.Register(MustNewLexer(
	&Config{
		Name:      "Befunge",
		Aliases:   []string{"befunge"},
		Filenames: []string{"*.befunge"},
		MimeTypes: []string{"application/x-befunge"},
	},
	Rules{
		"root": {},
	},
))
