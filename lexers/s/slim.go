package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Slim lexer.
var Slim = internal.Register(MustNewLexer(
	&Config{
		Name:      "Slim",
		Aliases:   []string{"slim"},
		Filenames: []string{"*.slim"},
		MimeTypes: []string{"text/x-slim"},
	},
	Rules{
		"root": {},
	},
))
