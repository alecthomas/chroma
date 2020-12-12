package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ezhil lexer.
var Ezhil = internal.Register(MustNewLexer(
	&Config{
		Name:      "Ezhil",
		Aliases:   []string{"ezhil"},
		Filenames: []string{"*.n"},
		MimeTypes: []string{"text/x-ezhil"},
	},
	Rules{
		"root": {},
	},
))
