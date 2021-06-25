package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MAQL lexer.
var MAQL = internal.Register(MustNewLexer(
	&Config{
		Name:      "MAQL",
		Aliases:   []string{"maql"},
		Filenames: []string{"*.maql"},
		MimeTypes: []string{"text/x-gooddata-maql", "application/x-gooddata-maql"},
	},
	Rules{
		"root": {},
	},
))
