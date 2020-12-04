package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// OpenEdge ABL lexer.
var OpenEdgeABL = internal.Register(MustNewLexer(
	&Config{
		Name:      "OpenEdge ABL",
		Aliases:   []string{"openedge", "abl", "progress"},
		Filenames: []string{"*.p", "*.cls"},
		MimeTypes: []string{"text/x-openedge", "application/x-openedge"},
	},
	Rules{
		"root": {},
	},
))
