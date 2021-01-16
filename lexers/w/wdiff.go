package w

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// WDiff lexer.
var WDiff = internal.Register(MustNewLexer(
	&Config{
		Name:      "WDiff",
		Aliases:   []string{"wdiff"},
		Filenames: []string{"*.wdiff"},
	},
	Rules{
		"root": {},
	},
))
