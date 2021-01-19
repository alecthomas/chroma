package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Logos lexer.
var Logos = internal.Register(MustNewLexer(
	&Config{
		Name:      "Logos",
		Aliases:   []string{"logos"},
		Filenames: []string{"*.x", "*.xi", "*.xm", "*.xmi"},
		MimeTypes: []string{"text/x-logos"},
		Priority:  0.25,
	},
	Rules{
		"root": {},
	},
))
