package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Inform6Template lexer.
var Inform6Template = internal.Register(MustNewLexer(
	&Config{
		Name:      "Inform 6 template",
		Aliases:   []string{"i6t"},
		Filenames: []string{"*.i6t"},
	},
	Rules{
		"root": {},
	},
))
