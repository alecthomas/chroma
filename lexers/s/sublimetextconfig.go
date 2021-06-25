package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SublimeTextConfig lexer.
var SublimeTextConfig = internal.Register(MustNewLexer(
	&Config{
		Name:      "Sublime Text Config",
		Aliases:   []string{"sublime"},
		Filenames: []string{"*.sublime-settings"},
	},
	Rules{
		"root": {},
	},
))
