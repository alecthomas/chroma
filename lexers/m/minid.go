package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MiniD lexer.
var MiniD = internal.Register(MustNewLexer(
	&Config{
		Name:    "MiniD",
		Aliases: []string{"minid"},
		// Don't lex .md as MiniD, reserve for Markdown.
		Filenames: []string{},
		MimeTypes: []string{"text/x-minidsrc"},
	},
	Rules{
		"root": {},
	},
))
