package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// FStar lexer.
var FStar = internal.Register(MustNewLexer(
	&Config{
		Name:      "FStar",
		Aliases:   []string{"fstar"},
		Filenames: []string{"*.fst", "*.fsti"},
		MimeTypes: []string{"text/x-fstar"},
	},
	Rules{
		"root": {},
	},
))
