package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// BbCode lexer.
var BbCode = internal.Register(MustNewLexer(
	&Config{
		Name:      "BBCode",
		Aliases:   []string{"bbcode"},
		MimeTypes: []string{"text/x-bbcode"},
	},
	Rules{
		"root": {},
	},
))
