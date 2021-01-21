package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Rd lexer. Lexer for R documentation (Rd) files.
var Rd = internal.Register(MustNewLexer(
	&Config{
		Name:      "Rd",
		Aliases:   []string{"rd"},
		Filenames: []string{"*.Rd"},
		MimeTypes: []string{"text/x-r-doc"},
	},
	Rules{
		"root": {},
	},
))
