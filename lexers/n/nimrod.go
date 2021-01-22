package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nimrod lexer.
var Nimrod = internal.Register(MustNewLexer(
	&Config{
		Name:      "Nimrod",
		Aliases:   []string{"nim", "nimrod"},
		Filenames: []string{"*.nim", "*.nimrod"},
		MimeTypes: []string{"text/x-nim"},
	},
	Rules{
		"root": {},
	},
))
