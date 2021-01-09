package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// DylanConsole lexer.
var DylanConsole = internal.Register(MustNewLexer(
	&Config{
		Name:      "Dylan session",
		Aliases:   []string{"dylan-console", "dylan-repl"},
		Filenames: []string{"*.dylan-console"},
		MimeTypes: []string{"text/x-dylan-console"},
	},
	Rules{
		"root": {},
	},
))
