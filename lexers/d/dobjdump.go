package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// DObjdump lexer.
var DObjdump = internal.Register(MustNewLexer(
	&Config{
		Name:      "d-objdump",
		Aliases:   []string{"d-objdump"},
		Filenames: []string{"*.d-objdump"},
		MimeTypes: []string{"text/x-d-objdump"},
	},
	Rules{
		"root": {},
	},
))
