package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Modelica lexer.
var Modelica = internal.Register(MustNewLexer(
	&Config{
		Name:      "Modelica",
		Aliases:   []string{"modelica"},
		Filenames: []string{"*.mo"},
		MimeTypes: []string{"text/x-modelica"},
	},
	Rules{
		"root": {},
	},
))
