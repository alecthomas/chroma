package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PowerShellSession lexer.
var PowerShellSession = internal.Register(MustNewLexer(
	&Config{
		Name:    "PowerShell Session",
		Aliases: []string{"ps1con"},
	},
	Rules{
		"root": {},
	},
))
