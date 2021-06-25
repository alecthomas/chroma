package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pike lexer.
var Pike = internal.Register(MustNewLexer(
	&Config{
		Name:      "Pike",
		Aliases:   []string{"pike"},
		Filenames: []string{"*.pike", "*.pmod"},
		MimeTypes: []string{"text/x-pike"},
	},
	Rules{
		"root": {},
	},
))
