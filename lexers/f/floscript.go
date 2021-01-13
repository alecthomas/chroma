package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// FloScript lexer.
var FloScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "FloScript",
		Aliases:   []string{"floscript", "flo"},
		Filenames: []string{"*.flo"},
	},
	Rules{
		"root": {},
	},
))
