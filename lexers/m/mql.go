package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MQL lexer.
var MQL = internal.Register(MustNewLexer(
	&Config{
		Name:      "MQL",
		Aliases:   []string{"mql", "mq4", "mq5", "mql4", "mql5"},
		Filenames: []string{"*.mq4", "*.mq5", "*.mqh"},
		MimeTypes: []string{"text/x-mql"},
	},
	Rules{
		"root": {},
	},
))
