package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// JuliaConsole lexer.
var JuliaConsole = internal.Register(MustNewLexer(
	&Config{
		Name:    "Julia console",
		Aliases: []string{"jlcon"},
	},
	Rules{
		"root": {},
	},
))
