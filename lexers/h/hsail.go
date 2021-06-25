package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Hsail lexer.
var Hsail = internal.Register(MustNewLexer(
	&Config{
		Name:      "HSAIL",
		Aliases:   []string{"hsail", "hsa"},
		Filenames: []string{"*.hsail"},
		MimeTypes: []string{"text/x-hsail"},
	},
	Rules{
		"root": {},
	},
))
