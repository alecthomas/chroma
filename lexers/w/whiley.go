package w

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Whiley lexer.
var Whiley = internal.Register(MustNewLexer(
	&Config{
		Name:      "Whiley",
		Aliases:   []string{"whiley"},
		Filenames: []string{"*.whiley"},
		MimeTypes: []string{"text/x-whiley"},
	},
	Rules{
		"root": {},
	},
))
