package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// GAP lexer.
var Gap = internal.Register(MustNewLexer(
	&Config{
		Name:      "GAP",
		Aliases:   []string{"gap"},
		Filenames: []string{"*.g", "*.gd", "*.gi", "*.gap"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
