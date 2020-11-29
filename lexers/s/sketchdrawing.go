package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SketchDrawing lexer.
var SketchDrawing = internal.Register(MustNewLexer(
	&Config{
		Name:      "Sketch Drawing",
		Aliases:   []string{"sketch"},
		Filenames: []string{"*.sketch"},
	},
	Rules{
		"root": {},
	},
))
