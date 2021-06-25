package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// DylanLid lexer.
var DylanLid = internal.Register(MustNewLexer(
	&Config{
		Name:      "DylanLID",
		Aliases:   []string{"dylan-lid", "lid"},
		Filenames: []string{"*.lid", "*.hdp"},
		MimeTypes: []string{"text/x-dylan-lid"},
	},
	Rules{
		"root": {},
	},
))
