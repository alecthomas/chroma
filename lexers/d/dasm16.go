package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Dasm16 lexer.
var Dasm16 = internal.Register(MustNewLexer(
	&Config{
		Name:      "DASM16",
		Aliases:   []string{"dasm16"},
		Filenames: []string{"*.dasm16", "*.dasm"},
		MimeTypes: []string{"text/x-dasm16"},
	},
	Rules{
		"root": {},
	},
))
