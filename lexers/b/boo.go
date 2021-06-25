package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Boo lexer.
var Boo = internal.Register(MustNewLexer(
	&Config{
		Name:      "Boo",
		Aliases:   []string{"boo"},
		Filenames: []string{"*.boo"},
		MimeTypes: []string{"text/x-boo"},
	},
	Rules{
		"root": {},
	},
))
