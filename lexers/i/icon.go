package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Icon lexer.
var Icon = internal.Register(MustNewLexer(
	&Config{
		Name:      "Icon",
		Aliases:   []string{"icon"},
		Filenames: []string{"*.icon", "*.ICON"},
	},
	Rules{
		"root": {},
	},
))
