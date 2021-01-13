package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Fantom lexer.
var Fantom = internal.Register(MustNewLexer(
	&Config{
		Name:      "Fantom",
		Aliases:   []string{"fan"},
		Filenames: []string{"*.fan"},
		MimeTypes: []string{"application/x-fantom"},
	},
	Rules{
		"root": {},
	},
))
