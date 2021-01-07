package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Bst lexer.
var Bst = internal.Register(MustNewLexer(
	&Config{
		Name:      "BST",
		Aliases:   []string{"bst", "bst-pybtex"},
		Filenames: []string{"*.bst"},
	},
	Rules{
		"root": {},
	},
))
