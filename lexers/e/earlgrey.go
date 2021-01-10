package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// EarlGrey lexer.
var EarlGrey = internal.Register(MustNewLexer(
	&Config{
		Name:      "Earl Grey",
		Aliases:   []string{"earl-grey", "earlgrey", "eg"},
		Filenames: []string{"*.eg"},
		MimeTypes: []string{"text/x-earl-grey"},
	},
	Rules{
		"root": {},
	},
))
