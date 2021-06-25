package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SARL lexer. For SARL <http://www.sarl.io> source code.
var SARL = internal.Register(MustNewLexer(
	&Config{
		Name:      "SARL",
		Aliases:   []string{"sarl"},
		Filenames: []string{"*.sarl"},
		MimeTypes: []string{"text/x-sarl"},
	},
	Rules{
		"root": {},
	},
))
