package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// AspectJ lexer.
var AspectJ = internal.Register(MustNewLexer(
	&Config{
		Name:      "AspectJ",
		Aliases:   []string{"aspectj"},
		Filenames: []string{"*.aj"},
		MimeTypes: []string{"text/x-aspectj"},
	},
	Rules{
		"root": {},
	},
))
