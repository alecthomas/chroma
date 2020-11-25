package i

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// IDL lexer.
var IDL = internal.Register(MustNewLexer(
	&Config{
		Name:      "IDL",
		Aliases:   []string{"idl"},
		Filenames: []string{"*.pro"},
		MimeTypes: []string{"text/idl"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// endelse seems to be unique to IDL, endswitch is rare at least.
	var result float32

	if strings.Contains(text, "endelse") {
		result += 0.2
	}

	if strings.Contains(text, "endswitch") {
		result += 0.01
	}

	return result
}))
