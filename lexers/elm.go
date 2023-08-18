package lexers

import (
	. "github.com/alecthomas/chroma/v2" // nolint
)

// Elm lexer.
var Elm = Register(MustNewXMLLexer(
	embedded,
	"embedded/elm.xml",
).SetConfig(
	&Config{
		Name:      "Elm",
		Aliases:   []string{"elm"},
		Filenames: []string{"*.elm"},
		MimeTypes: []string{"text/x-elm"},
	},
))
