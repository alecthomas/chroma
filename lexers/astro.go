package lexers

import (
	. "github.com/alecthomas/chroma/v3" // nolint
)

// Astro lexer.
var Astro = Register(DelegatingLexer(
	HTML,
	mustEmbeddedXMLLexer("embedded/astro.xml"),
))
