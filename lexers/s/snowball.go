package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Snowball lexer. Lexer for Snowball <http://snowballstem.org/> source code.
var Snowball = internal.Register(MustNewLexer(
	&Config{
		Name:      "Snowball",
		Aliases:   []string{"snowball"},
		Filenames: []string{"*.sbl"},
	},
	Rules{
		"root": {},
	},
))
