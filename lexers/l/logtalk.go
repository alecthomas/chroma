package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Logtalk lexer.
var Logtalk = internal.Register(MustNewLexer(
	&Config{
		Name:      "Logtalk",
		Aliases:   []string{"logtalk"},
		Filenames: []string{"*.lgt", "*.logtalk"},
		MimeTypes: []string{"text/x-logtalk"},
	},
	Rules{
		"root": {},
	},
))
