package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MIME lexer.
var MIME = internal.Register(MustNewLexer(
	&Config{
		Name:      "MIME",
		Aliases:   []string{"mime"},
		MimeTypes: []string{"multipart/mixed", "multipart/related", "multipart/alternative"},
	},
	Rules{
		"root": {},
	},
))
