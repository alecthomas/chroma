package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ParaSail lexer.
var ParaSail = internal.Register(MustNewLexer(
	&Config{
		Name:      "ParaSail",
		Aliases:   []string{"parasail"},
		Filenames: []string{"*.psi", "*.psl"},
		MimeTypes: []string{"text/x-parasail"},
	},
	Rules{
		"root": {},
	},
))
