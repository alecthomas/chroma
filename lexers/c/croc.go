package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Croc lexer.
var Croc = internal.Register(MustNewLexer(
	&Config{
		Name:      "Croc",
		Aliases:   []string{"croc"},
		Filenames: []string{"*.croc"},
		MimeTypes: []string{"text/x-crocsrc"},
	},
	Rules{
		"root": {},
	},
))
