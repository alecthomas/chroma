package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Sieve lexer. Lexer for sieve format.
var Sieve = internal.Register(MustNewLexer(
	&Config{
		Name:      "Sieve",
		Aliases:   []string{"sieve"},
		Filenames: []string{"*.siv", "*.sieve"},
	},
	Rules{
		"root": {},
	},
))
