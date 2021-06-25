package z

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Zephir lexer.
var Zephir = internal.Register(MustNewLexer(
	&Config{
		Name:      "Zephir",
		Aliases:   []string{"zephir"},
		Filenames: []string{"*.zep"},
	},
	Rules{
		"root": {},
	},
))
