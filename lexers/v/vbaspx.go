package v

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	vbAspxAnalyzerPageLanguageRe   = regexp.MustCompile(`(?i)Page\s*Language="Vb"`)
	vbAspxAnalyzerScriptLanguageRe = regexp.MustCompile(`(?i)script[^>]+language=["\']vb`)
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
).SetAnalyser(func(text string) float32 {
	if vbAspxAnalyzerPageLanguageRe.MatchString(text) {
		return 0.2
	}

	if vbAspxAnalyzerScriptLanguageRe.MatchString(text) {
		return 0.15
	}

	return 0
}))
