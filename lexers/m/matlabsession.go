package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MatlabSession lexer.
var MatlabSession = internal.Register(MustNewLexer(
	&Config{
		Name:    "Matlab session",
		Aliases: []string{"matlabsession"},
	},
	Rules{
		"root": {},
	},
))
