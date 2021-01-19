package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LiterateHaskell lexer.
var LiterateHaskell = internal.Register(MustNewLexer(
	&Config{
		Name:      "Literate Haskell",
		Aliases:   []string{"lhs", "literate-haskell", "lhaskell"},
		Filenames: []string{"*.lhs"},
		MimeTypes: []string{"text/x-literate-haskell"},
	},
	Rules{
		"root": {},
	},
))
