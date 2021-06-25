package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Treetop lexer. A lexer for Treetop <http://treetop.rubyforge.org/> grammars.
var Treetop = internal.Register(MustNewLexer(
	&Config{
		Name:      "Treetop",
		Aliases:   []string{"treetop"},
		Filenames: []string{"*.treetop", "*.tt"},
	},
	Rules{
		"root": {},
	},
))
