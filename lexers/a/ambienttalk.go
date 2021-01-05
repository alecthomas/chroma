package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// AmbientTalk lexer.
var AmbientTalk = internal.Register(MustNewLexer(
	&Config{
		Name:      "AmbientTalk",
		Aliases:   []string{"at", "ambienttalk", "ambienttalk/2"},
		Filenames: []string{"*.at"},
		MimeTypes: []string{"text/x-ambienttalk"},
	},
	Rules{
		"root": {},
	},
))
