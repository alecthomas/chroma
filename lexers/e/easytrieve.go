package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Easytrieve lexer.
var Easytrieve = internal.Register(MustNewLexer(
	&Config{
		Name:      "Easytrieve",
		Aliases:   []string{"easytrieve"},
		Filenames: []string{"*.ezt", "*.mac"},
		MimeTypes: []string{"text/x-easytrieve"},
	},
	Rules{
		"root": {},
	},
))
