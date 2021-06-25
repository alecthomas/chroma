package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TNT lexer. Lexer for Typographic Number Theory, as described in the book
// Gödel, Escher, Bach, by Douglas R. Hofstadter, or as summarized here:
// https://github.com/Kenny2github/language-tnt/blob/master/README.md#summary-of-tnt
var TNT = internal.Register(MustNewLexer(
	&Config{
		Name:      "Typographic Number Theory",
		Aliases:   []string{"tnt"},
		Filenames: []string{"*.tnt"},
	},
	Rules{
		"root": {},
	},
))
