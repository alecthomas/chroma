package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Monkey lexer.
var Monkey = internal.Register(MustNewLexer(
	&Config{
		Name:      "Monkey",
		Aliases:   []string{"monkey"},
		Filenames: []string{"*.monkey"},
		MimeTypes: []string{"text/x-monkey"},
	},
	Rules{
		"root": {},
	},
))
