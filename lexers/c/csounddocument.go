package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CsoundDocument lexer.
var CsoundDocument = internal.Register(MustNewLexer(
	&Config{
		Name:      "Csound Document",
		Aliases:   []string{"csound-document", "csound-csd"},
		Filenames: []string{"*.csd"},
	},
	Rules{
		"root": {},
	},
))
