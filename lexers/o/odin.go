package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ODIN lexer.
var ODIN = internal.Register(MustNewLexer(
	&Config{
		Name:      "ODIN",
		Aliases:   []string{"odin"},
		Filenames: []string{"*.odin"},
		MimeTypes: []string{"text/odin"},
	},
	Rules{
		"root": {},
	},
))
