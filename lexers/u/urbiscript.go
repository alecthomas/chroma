package u

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// UrbiScript lexer.
var UrbiScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "UrbiScript",
		Aliases:   []string{"urbiscript"},
		Filenames: []string{"*.u"},
		MimeTypes: []string{"application/x-urbiscript"},
	},
	Rules{
		"root": {},
	},
))
