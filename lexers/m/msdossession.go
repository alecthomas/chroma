package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MSDOSSession lexer.
var MSDOSSession = internal.Register(MustNewLexer(
	&Config{
		Name:    "MSDOS Session",
		Aliases: []string{"doscon"},
	},
	Rules{
		"root": {},
	},
))
