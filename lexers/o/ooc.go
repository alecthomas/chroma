package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ooc lexer.
var Ooc = internal.Register(MustNewLexer(
	&Config{
		Name:      "Ooc",
		Aliases:   []string{"ooc"},
		Filenames: []string{"*.ooc"},
		MimeTypes: []string{"text/x-ooc"},
	},
	Rules{
		"root": {},
	},
))
