package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Haml lexer.
var Haml = internal.Register(MustNewLexer(
	&Config{
		Name:      "Haml",
		Aliases:   []string{"haml"},
		Filenames: []string{"*.haml"},
		MimeTypes: []string{"text/x-haml"},
	},
	Rules{
		"root": {},
	},
))
