package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pug lexer.
var Pug = internal.Register(MustNewLexer(
	&Config{
		Name:      "Pug",
		Aliases:   []string{"pug", "jade"},
		Filenames: []string{"*.pug", "*.jade"},
		MimeTypes: []string{"text/x-pug", "text/x-jade"},
	},
	Rules{
		"root": {},
	},
))
