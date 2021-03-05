package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ReScript lexer.
var ReScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "ReScript",
		Aliases:   []string{"rescript"},
		Filenames: []string{"*.res", "*.resi"},
		MimeTypes: []string{"text/x-rescript"},
	},
	Rules{
		"root": {},
	},
))
