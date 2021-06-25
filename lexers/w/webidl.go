package w

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// WebIDL lexer.
var WebIDL = internal.Register(MustNewLexer(
	&Config{
		Name:      "Web IDL",
		Aliases:   []string{"webidl"},
		Filenames: []string{"*.webidl"},
	},
	Rules{
		"root": {},
	},
))
