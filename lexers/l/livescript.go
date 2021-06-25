package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LiveScript lexer.
var LiveScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "LiveScript",
		Aliases:   []string{"live-script", "livescript"},
		Filenames: []string{"*.ls"},
		MimeTypes: []string{"text/livescript"},
	},
	Rules{
		"root": {},
	},
))
