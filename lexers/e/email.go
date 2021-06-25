package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Email lexer.
var Email = internal.Register(MustNewLexer(
	&Config{
		Name:      "E-mail",
		Aliases:   []string{"email", "eml"},
		Filenames: []string{"*.eml"},
		MimeTypes: []string{"message/rfc822"},
	},
	Rules{
		"root": {},
	},
))
