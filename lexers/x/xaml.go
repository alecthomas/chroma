package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// XAML lexer.
var XAML = internal.Register(MustNewLexer(
	&Config{
		Name:      "XAML",
		Aliases:   []string{"xaml"},
		Filenames: []string{"*.xaml"},
		MimeTypes: []string{"application/xaml+xml"},
	},
	Rules{
		"root": {},
	},
))
