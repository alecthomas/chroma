package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Gettext lexer.
var Gettext = internal.Register(MustNewLexer(
	&Config{
		Name:      "Gettext Catalog",
		Aliases:   []string{"pot", "po"},
		Filenames: []string{"*.pot", "*.po"},
		MimeTypes: []string{"application/x-gettext", "text/x-gettext", "text/gettext"},
	},
	Rules{
		"root": {},
	},
))
