package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cryptol lexer.
var Cryptol = internal.Register(MustNewLexer(
	&Config{
		Name:      "Cryptol",
		Aliases:   []string{"cryptol", "cry"},
		Filenames: []string{"*.cry"},
		MimeTypes: []string{"text/x-cryptol"},
	},
	Rules{
		"root": {},
	},
))
