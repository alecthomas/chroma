package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Dg lexer.
var Dg = internal.Register(MustNewLexer(
	&Config{
		Name:      "dg",
		Aliases:   []string{"dg"},
		Filenames: []string{"*.dg"},
		MimeTypes: []string{"text/x-dg"},
	},
	Rules{
		"root": {},
	},
))
