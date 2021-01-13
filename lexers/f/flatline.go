package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Flatline lexer.
var Flatline = internal.Register(MustNewLexer(
	&Config{
		Name:      "Flatline",
		Aliases:   []string{"flatline"},
		MimeTypes: []string{"text/x-flatline"},
	},
	Rules{
		"root": {},
	},
))
