package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MOOCode lexer.
var MOOCode = internal.Register(MustNewLexer(
	&Config{
		Name:      "MOOCode",
		Aliases:   []string{"moocode", "moo"},
		Filenames: []string{"*.moo"},
		MimeTypes: []string{"text/x-moocode"},
	},
	Rules{
		"root": {},
	},
))
