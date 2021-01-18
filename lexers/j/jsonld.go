package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// JSONLd lexer.
var JSONLd = internal.Register(MustNewLexer(
	&Config{
		Name:      "JSON-LD",
		Aliases:   []string{"jsonld", "json-ld"},
		Filenames: []string{"*.jsonld"},
		MimeTypes: []string{"application/ld+json"},
	},
	Rules{
		"root": {},
	},
))
