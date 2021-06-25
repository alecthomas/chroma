package v

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VCL lexer.
var VCL = internal.Register(MustNewLexer(
	&Config{
		Name:      "VCL",
		Aliases:   []string{"vcl"},
		Filenames: []string{"*.vcl"},
		MimeTypes: []string{"text/x-vclsrc"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// If the very first line is 'vcl 4.0;' it's pretty much guaranteed
	// that this is VCL
	if strings.HasPrefix(text, "vcl 4.0;") {
		return 1.0
	}

	if len(text) > 1000 {
		text = text[:1000]
	}

	// Skip over comments and blank lines
	// This is accurate enough that returning 0.9 is reasonable.
	// Almost no VCL files start without some comments.
	if strings.Contains(text, "\nvcl 4.0;") {
		return 0.9
	}

	return 0
}))
