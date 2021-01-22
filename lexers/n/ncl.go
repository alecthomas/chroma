package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// NCL lexer.
var NCL = internal.Register(MustNewLexer(
	&Config{
		Name:      "NCL",
		Aliases:   []string{"ncl"},
		Filenames: []string{"*.ncl"},
		MimeTypes: []string{"text/ncl"},
	},
	Rules{
		"root": {},
	},
))
