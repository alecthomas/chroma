package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Felix lexer.
var Felix = internal.Register(MustNewLexer(
	&Config{
		Name:      "Felix",
		Aliases:   []string{"felix", "flx"},
		Filenames: []string{"*.flx", "*.flxh"},
		MimeTypes: []string{"text/x-felix"},
	},
	Rules{
		"root": {},
	},
))
