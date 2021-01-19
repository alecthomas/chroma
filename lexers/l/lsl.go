package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LSL lexer.
var LSL = internal.Register(MustNewLexer(
	&Config{
		Name:      "LSL",
		Aliases:   []string{"lsl"},
		Filenames: []string{"*.lsl"},
		MimeTypes: []string{"text/x-lsl"},
	},
	Rules{
		"root": {},
	},
))
