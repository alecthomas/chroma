package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// GoodDataCl lexer.
var GoodDataCl = internal.Register(MustNewLexer(
	&Config{
		Name:      "GoodData-CL",
		Aliases:   []string{"gooddata-cl"},
		Filenames: []string{"*.gdc"},
		MimeTypes: []string{"text/x-gooddata-cl"},
	},
	Rules{
		"root": {},
	},
))
