package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TcshSession lexer. Lexer for Tcsh sessions, i.e. command lines, including a
// prompt, interspersed with output.
var TcshSession = internal.Register(MustNewLexer(
	&Config{
		Name:    "Tcsh Session",
		Aliases: []string{"tcshcon"},
	},
	Rules{
		"root": {},
	},
))
