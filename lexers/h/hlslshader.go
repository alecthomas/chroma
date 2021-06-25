package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// HlslShader lexer.
var HlslShader = internal.Register(MustNewLexer(
	&Config{
		Name:      "HLSL",
		Aliases:   []string{"hsls"},
		Filenames: []string{"*.hlsl", "*.hlsli"},
		MimeTypes: []string{"text/x-hlsl"},
	},
	Rules{
		"root": {},
	},
))
