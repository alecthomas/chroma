package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PEG lexer.
var PEG = internal.Register(MustNewLexer(
	&Config{
		Name:      "PEG",
		Aliases:   []string{"peg"},
		Filenames: []string{"*.peg"},
		MimeTypes: []string{"text/x-peg"},
	},
	Rules{
		"root": {},
	},
))
