package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cirru lexer.
var Cirru = internal.Register(MustNewLexer(
	&Config{
		Name:      "Cirru",
		Aliases:   []string{"cirru"},
		Filenames: []string{"*.cirru"},
		MimeTypes: []string{"text/x-cirru"},
	},
	Rules{
		"root": {},
	},
))
