package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Fancy lexer.
var Fancy = internal.Register(MustNewLexer(
	&Config{
		Name:      "Fancy",
		Aliases:   []string{"fancy", "fy"},
		Filenames: []string{"*.fy", "*.fancypack"},
		MimeTypes: []string{"text/x-fancysrc"},
	},
	Rules{
		"root": {},
	},
))
