package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// GosuTemplate lexer.
var GosuTemplate = internal.Register(MustNewLexer(
	&Config{
		Name:      "Gosu Template",
		Aliases:   []string{"gst"},
		Filenames: []string{"*.gst"},
		MimeTypes: []string{"text/x-gosu-template"},
	},
	Rules{
		"root": {},
	},
))
