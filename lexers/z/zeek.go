package z

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Zeek lexer.
var Zeek = internal.Register(MustNewLexer(
	&Config{
		Name:      "Zeek",
		Aliases:   []string{"zeek", "bro"},
		Filenames: []string{"*.zeek", "*.bro"},
	},
	Rules{
		"root": {},
	},
))
