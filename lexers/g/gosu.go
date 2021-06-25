package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Gosu lexer.
var Gosu = internal.Register(MustNewLexer(
	&Config{
		Name:      "Gosu",
		Aliases:   []string{"gosu"},
		Filenames: []string{"*.gs", "*.gsx", "*.gsp", "*.vark"},
		MimeTypes: []string{"text/x-gosu"},
	},
	Rules{
		"root": {},
	},
))
