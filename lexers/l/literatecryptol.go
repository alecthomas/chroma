package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LiterateCryptol lexer.
var LiterateCryptol = internal.Register(MustNewLexer(
	&Config{
		Name:      "Literate Cryptol",
		Aliases:   []string{"lcry", "literate-cryptol", "lcryptol"},
		Filenames: []string{"*.lcry"},
		MimeTypes: []string{"text/x-literate-cryptol"},
	},
	Rules{
		"root": {},
	},
))
