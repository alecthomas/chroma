package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Silver lexer. For Silver <https://bitbucket.org/viperproject/silver> source code.
var Silver = internal.Register(MustNewLexer(
	&Config{
		Name:      "Silver",
		Aliases:   []string{"silver"},
		Filenames: []string{"*.sil", "*.vpr"},
	},
	Rules{
		"root": {},
	},
))
