package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// NesC lexer.
var NesC = internal.Register(MustNewLexer(
	&Config{
		Name:      "nesC",
		Aliases:   []string{"nesc"},
		Filenames: []string{"*.nc"},
		MimeTypes: []string{"text/x-nescsrc"},
	},
	Rules{
		"root": {},
	},
))
