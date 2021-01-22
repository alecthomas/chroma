package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Objdump lexer.
var Objdump = internal.Register(MustNewLexer(
	&Config{
		Name:      "objdump",
		Aliases:   []string{"objdump"},
		Filenames: []string{"*.objdump"},
		MimeTypes: []string{"text/x-objdump"},
	},
	Rules{
		"root": {},
	},
))
