package t

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TADS 3 lexer.
var Tads3 = internal.Register(MustNewLexer(
	&Config{
		Name:      "TADS 3",
		Aliases:   []string{"tads3"},
		Filenames: []string{"*.t"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// This is a rather generic descriptive language without strong
	// identifiers. It looks like a 'GameMainDef' has to be present,
	// and/or a 'versionInfo' with an 'IFID' field.
	var result float32

	if strings.Contains(text, "__TADS") || strings.Contains(text, "GameMainDef") {
		result += 0.2
	}

	// This is a fairly unique keyword which is likely used in source as well.
	if strings.Contains(text, "versionInfo") && strings.Contains(text, "IFID") {
		result += 0.1
	}

	return result
}))
