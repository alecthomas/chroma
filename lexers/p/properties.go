package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Properties lexer.
var Properties = internal.Register(MustNewLexer(
	&Config{
		Name:      "Properties",
		Aliases:   []string{"properties", "jproperties"},
		Filenames: []string{"*.properties"},
		MimeTypes: []string{"text/x-java-properties"},
	},
	Rules{
		"root": {},
	},
))
