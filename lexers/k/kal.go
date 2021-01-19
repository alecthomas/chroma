package k

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Kal lexer.
var Kal = internal.Register(MustNewLexer(
	&Config{
		Name:      "Kal",
		Aliases:   []string{"kal"},
		Filenames: []string{"*.kal"},
		MimeTypes: []string{"text/kal", "application/kal"},
	},
	Rules{
		"root": {},
	},
))
