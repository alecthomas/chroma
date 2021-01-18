package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Jsp lexer.
var Jsp = internal.Register(MustNewLexer(
	&Config{
		Name:      "Java Server Page",
		Aliases:   []string{"jsp"},
		Filenames: []string{"*.jsp"},
		MimeTypes: []string{"application/x-jsp"},
	},
	Rules{
		"root": {},
	},
))
