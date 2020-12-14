package u

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ucode lexer.
var Ucode = internal.Register(MustNewLexer(
	&Config{
		Name:      "ucode",
		Aliases:   []string{"ucode"},
		Filenames: []string{"*.u", "*.u1", "*.u2"},
	},
	Rules{
		"root": {},
	},
))
