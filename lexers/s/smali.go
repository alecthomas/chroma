package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Smali lexer.
var Smali = internal.Register(MustNewLexer(
	&Config{
		Name:      "Smali",
		Aliases:   []string{"smali"},
		Filenames: []string{"*.smali"},
		MimeTypes: []string{"text/smali"},
	},
	Rules{
		"root": {},
	},
))
