package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Tera Term macro lexer.
var TeraTerm = internal.Register(MustNewLexer(
	&Config{
		Name:      "Tera Term macro",
		Aliases:   []string{"ttl", "teraterm", "teratermmacro"},
		Filenames: []string{"*.ttl"},
		MimeTypes: []string{"text/x-teratermmacro"},
	},
	Rules{
		"root": {},
	},
))
