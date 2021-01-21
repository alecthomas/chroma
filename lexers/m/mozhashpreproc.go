package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MozPreprocHash lexer.
var MozPreprocHash = internal.Register(MustNewLexer(
	&Config{
		Name:    "mozhashpreproc",
		Aliases: []string{"mozhashpreproc"},
	},
	Rules{
		"root": {},
	},
))
