package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CSharpAspx lexer.
var CSharpAspx = internal.Register(MustNewLexer(
	&Config{
		Name:      "aspx-cs",
		Aliases:   []string{"aspx-cs"},
		Filenames: []string{"*.aspx", "*.asax", "*.ascx", "*.ashx", "*.asmx", "*.axd"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
