package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Opa lexer.
var Opa = internal.Register(MustNewLexer(
	&Config{
		Name:      "Opa",
		Aliases:   []string{"opa"},
		Filenames: []string{"*.opa"},
		MimeTypes: []string{"text/x-opa"},
	},
	Rules{
		"root": {},
	},
))
