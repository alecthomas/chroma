package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Erb lexer.
var Erb = internal.Register(MustNewLexer(
	&Config{
		Name:      "ERB",
		Aliases:   []string{"erb"},
		MimeTypes: []string{"application/x-ruby-templating"},
	},
	Rules{
		"root": {},
	},
))
