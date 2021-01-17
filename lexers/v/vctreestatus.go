package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VCTreeStatus lexer.
var VCTreeStatus = internal.Register(MustNewLexer(
	&Config{
		Name:    "VCTreeStatus",
		Aliases: []string{"vctreestatus"},
	},
	Rules{
		"root": {},
	},
))
