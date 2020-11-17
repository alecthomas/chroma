package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Limbo lexer.
var Limbo = internal.Register(MustNewLexer(
	&Config{
		Name:      "Limbo",
		Aliases:   []string{"limbo"},
		Filenames: []string{"*.b"},
		MimeTypes: []string{"text/limbo"},
	},
	Rules{
		"root": {},
	},
))
