package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Crmsh lexer.
var Crmsh = internal.Register(MustNewLexer(
	&Config{
		Name:      "Crmsh",
		Aliases:   []string{"crmsh", "pcmk"},
		Filenames: []string{"*.crmsh", "*.pcmk"},
	},
	Rules{
		"root": {},
	},
))
