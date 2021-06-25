package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Shen lexer. Lexer for Shen <http://shenlanguage.org/> source code.
var Shen = internal.Register(MustNewLexer(
	&Config{
		Name:      "Shen",
		Aliases:   []string{"shen"},
		Filenames: []string{"*.shen"},
		MimeTypes: []string{"text/x-shen", "application/x-shen"},
	},
	Rules{
		"root": {},
	},
))
