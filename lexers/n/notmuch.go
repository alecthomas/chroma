package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Notmuch lexer.
var Notmuch = internal.Register(MustNewLexer(
	&Config{
		Name:    "Notmuch",
		Aliases: []string{"notmuch"},
	},
	Rules{
		"root": {},
	},
))
