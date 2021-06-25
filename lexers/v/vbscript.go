package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VBScript lexer.
var VBScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "VBScript",
		Aliases:   []string{"vbscript"},
		Filenames: []string{"*.vbs", "*.VBS"},
	},
	Rules{
		"root": {},
	},
))
