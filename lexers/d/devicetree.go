package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Devicetree lexer.
var Devicetree = internal.Register(MustNewLexer(
	&Config{
		Name:      "Devicetree",
		Aliases:   []string{"devicetree", "dts"},
		Filenames: []string{"*.dts", "*.dtsi"},
		MimeTypes: []string{"text/x-c"},
	},
	Rules{
		"root": {},
	},
))
