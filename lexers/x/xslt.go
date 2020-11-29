package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// XSLT lexer.
var XSLT = internal.Register(MustNewLexer(
	&Config{
		Name:    "XSLT",
		Aliases: []string{"xslt"},
		// xpl is XProc
		Filenames: []string{"*.xsl", "*.xslt", "*.xpl"},
		MimeTypes: []string{"application/xsl+xml", "application/xslt+xml"},
	},
	Rules{
		"root": {},
	},
))
