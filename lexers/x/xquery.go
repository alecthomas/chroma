package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// XQuery lexer.
var XQuery = internal.Register(MustNewLexer(
	&Config{
		Name:      "XQuery",
		Aliases:   []string{"xquery", "xqy", "xq", "xql", "xqm"},
		Filenames: []string{"*.xqy", "*.xquery", "*.xq", "*.xql", "*.xqm"},
		MimeTypes: []string{"text/xquery", "application/xquery"},
	},
	Rules{
		"root": {},
	},
))
