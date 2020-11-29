package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Lasso lexer.
var Lasso = internal.Register(MustNewLexer(
	&Config{
		Name:    "Lasso",
		Aliases: []string{"lasso", "lassoscript"},
		Filenames: []string{
			"*.lasso",
			"*.lasso[89]",
		},
		AliasFilenames: []string{
			"*.incl",
			"*.inc",
			"*.las",
		},
		MimeTypes: []string{"text/x-lasso"},
	},
	Rules{
		"root": {},
	},
))
