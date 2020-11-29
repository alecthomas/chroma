package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Delphi lexer.
var Delphi = internal.Register(MustNewLexer(
	&Config{
		Name:      "Delphi",
		Aliases:   []string{"delphi", "pas", "pascal", "objectpascal"},
		Filenames: []string{"*.pas", "*.dpr"},
		MimeTypes: []string{"text/x-pascal"},
	},
	Rules{
		"root": {},
	},
))
