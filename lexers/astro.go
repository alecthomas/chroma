package lexers

import (
	. "github.com/alecthomas/chroma/v3" // nolint
)

// Astro lexer.
var Astro = Register(DelegatingLexer(
	HTML,
	MustNewXMLLexer(embedded, "embedded/astro.xml"),
))
