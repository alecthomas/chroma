package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LiterateIdris lexer.
var LiterateIdris = internal.Register(MustNewLexer(
	&Config{
		Name:      "Literate Idris",
		Aliases:   []string{"lidr", "literate-idris", "lidris"},
		Filenames: []string{"*.lidr"},
		MimeTypes: []string{"text/x-literate-idris"},
	},
	Rules{
		"root": {},
	},
))
