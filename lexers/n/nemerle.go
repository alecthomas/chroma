package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nemerle lexer.
var Nemerle = internal.Register(MustNewLexer(
	&Config{
		Name:      "Nemerle",
		Aliases:   []string{"nemerle"},
		Filenames: []string{"*.n"},
		// inferred
		MimeTypes: []string{"text/x-nemerle"},
	},
	Rules{
		"root": {},
	},
))
