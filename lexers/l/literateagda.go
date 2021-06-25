package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LiterateAgda lexer.
var LiterateAgda = internal.Register(MustNewLexer(
	&Config{
		Name:      "Literate Agda",
		Aliases:   []string{"lagda", "literate-agda"},
		Filenames: []string{"*.lagda"},
		MimeTypes: []string{"text/x-literate-agda"},
	},
	Rules{
		"root": {},
	},
))
