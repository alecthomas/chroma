package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// X10 lexer.
var X10 = internal.Register(MustNewLexer(
	&Config{
		Name:      "X10",
		Aliases:   []string{"x10", "xten"},
		Filenames: []string{"*.x10"},
		MimeTypes: []string{"text/x-x10"},
	},
	Rules{
		"root": {},
	},
))
