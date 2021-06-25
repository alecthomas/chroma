package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// AutoIt lexer.
var AutoIt = internal.Register(MustNewLexer(
	&Config{
		Name:      "AutoIt",
		Aliases:   []string{"autoit"},
		Filenames: []string{"*.au3"},
		MimeTypes: []string{"text/x-autoit"},
	},
	Rules{
		"root": {},
	},
))
