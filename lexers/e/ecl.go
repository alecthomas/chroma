package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ECL lexer.
var Ecl = internal.Register(MustNewLexer(
	&Config{
		Name:      "ECL",
		Aliases:   []string{"ecl"},
		Filenames: []string{"*.ecl"},
		MimeTypes: []string{"application/x-ecl"},
	},
	Rules{
		"root": {},
	},
))
