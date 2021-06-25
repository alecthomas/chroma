package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Bare lexer.
var Bare = internal.Register(MustNewLexer(
	&Config{
		Name:      "BARE",
		Aliases:   []string{"bare"},
		Filenames: []string{"*.bare"},
	},
	Rules{
		"root": {},
	},
))
