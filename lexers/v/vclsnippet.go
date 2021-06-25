package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VCLSnippets lexer.
var VCLSnippets = internal.Register(MustNewLexer(
	&Config{
		Name:      "VCLSnippets",
		Aliases:   []string{"vclsnippets", "vclsnippet"},
		MimeTypes: []string{"text/x-vclsnippet"},
	},
	Rules{
		"root": {},
	},
))
