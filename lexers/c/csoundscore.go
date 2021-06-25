package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CsoundScore lexer.
var CsoundScore = internal.Register(MustNewLexer(
	&Config{
		Name:      "Csound Score",
		Aliases:   []string{"csound-score", "csound-sco"},
		Filenames: []string{"*.sco"},
	},
	Rules{
		"root": {},
	},
))
