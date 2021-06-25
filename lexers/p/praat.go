package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Praat lexer.
var Praat = internal.Register(MustNewLexer(
	&Config{
		Name:      "Praat",
		Aliases:   []string{"praat"},
		Filenames: []string{"*.praat", "*.proc", "*.psc"},
	},
	Rules{
		"root": {},
	},
))
