package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Asymptote lexer.
var Asymptote = internal.Register(MustNewLexer(
	&Config{
		Name:      "Asymptote",
		Aliases:   []string{"asy", "asymptote"},
		Filenames: []string{"*.asy"},
		MimeTypes: []string{"text/x-asymptote"},
	},
	Rules{
		"root": {},
	},
))
