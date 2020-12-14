package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// scdoc lexer.
var Scdoc = internal.Register(MustNewLexer(
	&Config{
		Name:      "scdoc",
		Aliases:   []string{"scdoc", "scd"},
		Filenames: []string{"*.scd", "*.scdoc"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
