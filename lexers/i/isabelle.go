package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Isabelle lexer.
var Isabelle = internal.Register(MustNewLexer(
	&Config{
		Name:      "Isabelle",
		Aliases:   []string{"isabelle"},
		Filenames: []string{"*.thy"},
		MimeTypes: []string{"text/x-isabelle"},
	},
	Rules{
		"root": {},
	},
))
