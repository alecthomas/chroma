package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ioke lexer.
var Ioke = internal.Register(MustNewLexer(
	&Config{
		Name:      "Ioke",
		Aliases:   []string{"ioke", "ik"},
		Filenames: []string{"*.ik"},
		MimeTypes: []string{"text/x-iokesrc"},
	},
	Rules{
		"root": {},
	},
))
