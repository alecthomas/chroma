package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Lean lexer.
var Lean = internal.Register(MustNewLexer(
	&Config{
		Name:      "Lean",
		Aliases:   []string{"lean"},
		Filenames: []string{"*.lean"},
		MimeTypes: []string{"text/x-lean"},
	},
	Rules{
		"root": {},
	},
))
