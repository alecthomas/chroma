package c

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	csharpAspxAnalyzerPageLanguageRe   = regexp.MustCompile(`(?i)Page\s*Language="C#"`)
	csharpAspxAnalyzerScriptLanguageRe = regexp.MustCompile(`(?i)script[^>]+language=["\']C#`)
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
).SetAnalyser(func(text string) float32 {
	if csharpAspxAnalyzerPageLanguageRe.MatchString(text) {
		return 0.2
	}

	if csharpAspxAnalyzerScriptLanguageRe.MatchString(text) {
		return 0.15
	}

	return 0
}))
