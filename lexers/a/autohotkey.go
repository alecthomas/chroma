package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Autohotkey lexer.
var Autohotkey = internal.Register(MustNewLexer(
	&Config{
		Name:      "autohotkey",
		Aliases:   []string{"ahk", "autohotkey"},
		Filenames: []string{"*.ahk", "*.ahkl"},
		MimeTypes: []string{"text/x-autohotkey"},
	},
	Rules{
		"root": {},
	},
))
