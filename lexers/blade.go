package lexers

import (
	. "github.com/alecthomas/chroma/v3" // nolint
)

// Blade lexer is Laravel's Blade templating language embedded in HTML.
//
// blade.xml holds the rules and metadata; this wrapper exists only to delegate
// the text Blade leaves alone to the HTML lexer, which XML cannot express.
var Blade = Register(DelegatingLexer(HTML, MustNewXMLLexer(
	embedded,
	"embedded/blade.xml",
)))
