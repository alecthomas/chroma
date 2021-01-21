package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Mask lexer.
var Mask = internal.Register(MustNewLexer(
	&Config{
		Name:      "Mask",
		Aliases:   []string{"mask"},
		Filenames: []string{"*.mask"},
		MimeTypes: []string{"text/x-mask"},
	},
	Rules{
		"root": {},
	},
))
