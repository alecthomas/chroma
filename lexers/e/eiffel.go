package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Eiffel lexer.
var Eiffel = internal.Register(MustNewLexer(
	&Config{
		Name:      "Eiffel",
		Aliases:   []string{"eiffel"},
		Filenames: []string{"*.e"},
		MimeTypes: []string{"text/x-eiffel"},
	},
	Rules{
		"root": {},
	},
))
