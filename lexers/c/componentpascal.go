package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ComponentPascal lexer.
var ComponentPascal = internal.Register(MustNewLexer(
	&Config{
		Name:      "Component Pascal",
		Aliases:   []string{"componentpascal", "cp"},
		Filenames: []string{"*.cp", "*.cps"},
		MimeTypes: []string{"text/x-component-pascal"},
	},
	Rules{
		"root": {},
	},
))
