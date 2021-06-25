package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Mustache lexer.
var Mustache = internal.Register(MustNewLexer(
	&Config{
		Name:      "Mustache",
		Aliases:   []string{"mustache"},
		Filenames: []string{"*.mustache"},
		MimeTypes: []string{"text/x-mustache-template"},
	},
	Rules{
		"root": {},
	},
))
