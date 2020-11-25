package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// IDL lexer.
var Idl = internal.Register(MustNewLexer(
	&Config{
		Name:      "IDL",
		Aliases:   []string{"idl"},
		Filenames: []string{"*.pro"},
		MimeTypes: []string{"text/idl"},
	},
	Rules{
		"root": {},
	},
))
