package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RConsole lexer. For R console transcripts or R CMD BATCH output files.
var RConsole = internal.Register(MustNewLexer(
	&Config{
		Name:      "RConsole",
		Aliases:   []string{"rconsole", "rout"},
		Filenames: []string{"*.Rout"},
	},
	Rules{
		"root": {},
	},
))
