package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Mscgen lexer.
var Mscgen = internal.Register(MustNewLexer(
	&Config{
		Name:      "Mscgen",
		Aliases:   []string{"mscgen", "msc"},
		Filenames: []string{"*.msc"},
	},
	Rules{
		"root": {},
	},
))
