package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MozPreprocPercent lexer.
var MozPreprocPercent = internal.Register(MustNewLexer(
	&Config{
		Name:    "mozpercentpreproc",
		Aliases: []string{"mozpercentpreproc"},
	},
	Rules{
		"root": {},
	},
))
