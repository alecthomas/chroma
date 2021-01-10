package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ec lexer.
var Ec = internal.Register(MustNewLexer(
	&Config{
		Name:      "eC",
		Aliases:   []string{"ec"},
		Filenames: []string{"*.ec", "*.eh"},
		MimeTypes: []string{"text/x-echdr", "text/x-ecsrc"},
	},
	Rules{
		"root": {},
	},
))
