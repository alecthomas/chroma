package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ShExC lexer. Lexer for ShExC <https://shex.io/shex-semantics/#shexc> shape expressions language syntax.
var ShExC = internal.Register(MustNewLexer(
	&Config{
		Name:      "ShExC",
		Aliases:   []string{"shexc", "shex"},
		Filenames: []string{"*.shex"},
		MimeTypes: []string{"text/shex"},
	},
	Rules{
		"root": {},
	},
))
