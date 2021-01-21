package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Monte lexer.
var Monte = internal.Register(MustNewLexer(
	&Config{
		Name:      "Monte",
		Aliases:   []string{"monte"},
		Filenames: []string{"*.mt"},
	},
	Rules{
		"root": {},
	},
))
