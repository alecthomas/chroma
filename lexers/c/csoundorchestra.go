package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CsoundOrchestra lexer.
var CsoundOrchestra = internal.Register(MustNewLexer(
	&Config{
		Name:      "Csound Orchestra",
		Aliases:   []string{"csound", "csound-orc"},
		Filenames: []string{"*.orc", "*.udo"},
	},
	Rules{
		"root": {},
	},
))
