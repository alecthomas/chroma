package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MoonScript lexer.
var MoonScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "MoonScript",
		Aliases:   []string{"moon", "moonscript"},
		MimeTypes: []string{"text/x-moonscript", "application/x-moonscript"},
	},
	Rules{
		"root": {},
	},
))
