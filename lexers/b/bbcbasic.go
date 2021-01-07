package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// BBC Basic lexer.
var BbcBasic = internal.Register(MustNewLexer(
	&Config{
		Name:      "BBC Basic",
		Aliases:   []string{"bbcbasic"},
		Filenames: []string{"*.bbc"},
	},
	Rules{
		"root": {},
	},
))
