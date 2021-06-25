package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CObjdump lexer.
var CObjdump = internal.Register(MustNewLexer(
	&Config{
		Name:      "c-objdump",
		Aliases:   []string{"c-objdump"},
		Filenames: []string{"*.c-objdump"},
		MimeTypes: []string{"text/x-c-objdump"},
	},
	Rules{
		"root": {},
	},
))
