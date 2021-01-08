package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CAmkES lexer.
var CAmkES = internal.Register(MustNewLexer(
	&Config{
		Name:      "CAmkES",
		Aliases:   []string{"camkes", "idl4"},
		Filenames: []string{"*.camkes", "*.idl4"},
	},
	Rules{
		"root": {},
	},
))
