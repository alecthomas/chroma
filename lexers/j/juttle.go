package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Juttle lexer.
var Juttle = internal.Register(MustNewLexer(
	&Config{
		Name:      "Juttle",
		Aliases:   []string{"juttle"},
		Filenames: []string{"*.juttle"},
		MimeTypes: []string{"application/juttle", "application/x-juttle", "text/x-juttle", "text/juttle"},
	},
	Rules{
		"root": {},
	},
))
