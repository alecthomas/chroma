package u

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// USD lexer.
var USD = internal.Register(MustNewLexer(
	&Config{
		Name:      "USD",
		Aliases:   []string{"usd", "usda"},
		Filenames: []string{"*.usd", "*.usda"},
	},
	Rules{
		"root": {},
	},
))
