package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cypher lexer.
var Cypher = internal.Register(MustNewLexer(
	&Config{
		Name:      "Cypher",
		Aliases:   []string{"cypher"},
		Filenames: []string{"*.cyp", "*.cypher"},
	},
	Rules{
		"root": {},
	},
))
