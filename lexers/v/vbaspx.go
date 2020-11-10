package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VBNetAspx lexer.
var VBNetAspx = internal.Register(MustNewLexer(
	&Config{
		Name:      "aspx-vb",
		Aliases:   []string{"aspx-vb"},
		Filenames: []string{"*.aspx", "*.asax", "*.ascx", "*.ashx", "*.asmx", "*.axd"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
