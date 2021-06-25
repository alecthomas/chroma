package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LESS lexer.
var LESS = internal.Register(MustNewLexer(
	&Config{
		Name:      "LessCss",
		Aliases:   []string{"less"},
		Filenames: []string{"*.less"},
		MimeTypes: []string{"text/x-less-css"},
	},
	Rules{
		"root": {},
	},
))
