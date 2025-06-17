package lexers

import (
	"github.com/crowyy03/chroma/v2"
)

// HTML lexer.
var HTML = chroma.MustNewXMLLexer(embedded, "embedded/html.xml")
