package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// JAGS lexer.
var Jags = internal.Register(MustNewLexer(
	&Config{
		Name:      "JAGS",
		Aliases:   []string{"jags"},
		Filenames: []string{"*.jag", "*.bug"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
