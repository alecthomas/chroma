package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// DarcsPatch lexer.
var DarcsPatch = internal.Register(MustNewLexer(
	&Config{
		Name:      "Darcs Patch",
		Aliases:   []string{"dpatch"},
		Filenames: []string{"*.dpatch", "*.darcspatch"},
	},
	Rules{
		"root": {},
	},
))
