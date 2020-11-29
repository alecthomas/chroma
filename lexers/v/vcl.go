package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VCL lexer.
var VCL = internal.Register(MustNewLexer(
	&Config{
		Name:      "VCL",
		Aliases:   []string{"vcl"},
		Filenames: []string{"*.vcl"},
		MimeTypes: []string{"text/x-vclsrc"},
	},
	Rules{
		"root": {},
	},
))
