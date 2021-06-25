package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Freefem lexer.
var Freefem = internal.Register(MustNewLexer(
	&Config{
		Name:      "Freefem",
		Aliases:   []string{"freefem"},
		Filenames: []string{"*.edp"},
		MimeTypes: []string{"text/x-freefem"},
	},
	Rules{
		"root": {},
	},
))
