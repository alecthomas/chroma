package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// NewLisp lexer.
var NewLisp = internal.Register(MustNewLexer(
	&Config{
		Name:      "NewLisp",
		Aliases:   []string{"newlisp"},
		Filenames: []string{"*.lsp", "*.nl", "*.kif"},
		MimeTypes: []string{"text/x-newlisp", "application/x-newlisp"},
	},
	Rules{
		"root": {},
	},
))
