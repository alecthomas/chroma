package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MuPAD lexer.
var MuPAD = internal.Register(MustNewLexer(
	&Config{
		Name:      "MuPAD",
		Aliases:   []string{"mupad"},
		Filenames: []string{"*.mu"},
	},
	Rules{
		"root": {},
	},
))
