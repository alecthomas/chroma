package x

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/xml"
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
).SetAnalyser(func(text string) float32 {
	if xml.MatchString(text) && strings.Contains(text, "<xsl") {
		return 0.8
	}

	return 0
}))
