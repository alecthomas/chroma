package lexers

import (
	"github.com/alecthomas/chroma/v3"
)

// HTML lexer.
var HTML = chroma.MustNewXMLLexer(embedded, "embedded/html.xml")
