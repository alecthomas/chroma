package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MiniScript lexer.
var MiniScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "MiniScript",
		Aliases:   []string{"ms", "miniscript"},
		Filenames: []string{"*.ms"},
		MimeTypes: []string{"text/x-miniscript", "application/x-miniscript"},
	},
	Rules{
		"root": {},
	},
))
