package s

import (
	"regexp"
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/xml"
)

var sspAnalyserRe = regexp.MustCompile(`val \w+\s*:`)

// SSP lexer. Lexer for Scalate Server Pages.
var SSP = internal.Register(MustNewLexer(
	&Config{
		Name:      "Scalate Server Page",
		Aliases:   []string{"ssp"},
		Filenames: []string{"*.ssp"},
		MimeTypes: []string{"application/x-ssp"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	var result float64

	if sspAnalyserRe.MatchString(text) {
		result += 0.6
	}

	if xml.MatchString(text) {
		result += 0.2
	}

	if strings.Contains(text, "<%") && strings.Contains(text, "%>") {
		result += 0.1
	}

	return float32(result)
}))
