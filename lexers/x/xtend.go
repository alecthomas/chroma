package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Xtend lexer.
var Xtend = internal.Register(MustNewLexer(
	&Config{
		Name:      "Xtend",
		Aliases:   []string{"xtend"},
		Filenames: []string{"*.xtend"},
		MimeTypes: []string{"text/x-xtend"},
	},
	Rules{
		"root": {},
	},
))
