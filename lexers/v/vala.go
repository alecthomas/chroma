package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Vala lexer.
var Vala = internal.Register(MustNewLexer(
	&Config{
		Name:      "Vala",
		Aliases:   []string{"vala", "vapi"},
		Filenames: []string{"*.vala", "*.vapi"},
		MimeTypes: []string{"text/x-vala"},
	},
	Rules{
		"root": {},
	},
))
