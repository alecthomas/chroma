package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Red lexer. A Red-language <http://www.red-lang.org/> lexer.
var Red = internal.Register(MustNewLexer(
	&Config{
		Name:      "Red",
		Aliases:   []string{"red", "red/system"},
		Filenames: []string{"*.red", "*.reds"},
		MimeTypes: []string{"text/x-red", "text/x-red-system"},
	},
	Rules{
		"root": {},
	},
))
