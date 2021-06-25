package k

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Koka lexer.
var Koka = internal.Register(MustNewLexer(
	&Config{
		Name:      "Koka",
		Aliases:   []string{"koka"},
		Filenames: []string{"*.kk", "*.kki"},
		MimeTypes: []string{"text/x-koka"},
	},
	Rules{
		"root": {},
	},
))
