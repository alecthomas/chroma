package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SourcePawn lexer.
var SourcePawn = internal.Register(MustNewLexer(
	&Config{
		Name:      "SourcePawn",
		Aliases:   []string{"sp"},
		Filenames: []string{"*.sp"},
		MimeTypes: []string{"text/x-sourcepawn"},
	},
	Rules{
		"root": {},
	},
))
