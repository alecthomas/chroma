package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TAP lexer. For Test Anything Protocol (TAP) output.
var TAP = internal.Register(MustNewLexer(
	&Config{
		Name:      "TAP",
		Aliases:   []string{"tap"},
		Filenames: []string{"*.tap"},
	},
	Rules{
		"root": {},
	},
))
