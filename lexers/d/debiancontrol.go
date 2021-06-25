package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// DebianControl lexer.
var DebianControl = internal.Register(MustNewLexer(
	&Config{
		Name:      "Debian Control file",
		Aliases:   []string{"control", "debcontrol"},
		Filenames: []string{"control"},
	},
	Rules{
		"root": {},
	},
))
