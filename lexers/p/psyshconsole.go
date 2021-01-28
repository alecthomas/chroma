package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PsyshConsole lexer.
var PsyshConsole = internal.Register(MustNewLexer(
	&Config{
		Name:    "PsySH console session for PHP",
		Aliases: []string{"psysh"},
	},
	Rules{
		"root": {},
	},
))
