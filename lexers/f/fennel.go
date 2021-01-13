package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Fennel lexer.
var Fennel = internal.Register(MustNewLexer(
	&Config{
		Name:      "Fennel",
		Aliases:   []string{"fennel", "fnl"},
		Filenames: []string{"*.fnl"},
	},
	Rules{
		"root": {},
	},
))
