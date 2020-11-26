package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ca65 lexer.
var Ca65 = internal.Register(MustNewLexer(
	&Config{
		Name:      "ca65 assembler",
		Aliases:   []string{"ca65"},
		Filenames: []string{"*.s"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
