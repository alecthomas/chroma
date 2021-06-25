package g

import (
	"regexp"
	"unicode"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var groffAlphanumericRe = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// Groff lexer.
var Groff = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Groff",
		Aliases:   []string{"groff", "nroff", "man"},
		Filenames: []string{"*.[1-9]", "*.1p", "*.3pm", "*.man"},
		MimeTypes: []string{"application/x-troff", "text/troff"},
	},
	func() Rules {
		return Rules{
			"root": {
				{`(\.)(\w+)`, ByGroups(Text, Keyword), Push("request")},
				{`\.`, Punctuation, Push("request")},
				{`[^\\\n]+`, Text, Push("textline")},
				Default(Push("textline")),
			},
			"textline": {
				Include("escapes"),
				{`[^\\\n]+`, Text, nil},
				{`\n`, Text, Pop(1)},
			},
			"escapes": {
				{`\\"[^\n]*`, Comment, nil},
				{`\\[fn]\w`, LiteralStringEscape, nil},
				{`\\\(.{2}`, LiteralStringEscape, nil},
				{`\\.\[.*\]`, LiteralStringEscape, nil},
				{`\\.`, LiteralStringEscape, nil},
				{`\\\n`, Text, Push("request")},
			},
			"request": {
				{`\n`, Text, Pop(1)},
				Include("escapes"),
				{`"[^\n"]+"`, LiteralStringDouble, nil},
				{`\d+`, LiteralNumber, nil},
				{`\S+`, LiteralString, nil},
				{`\s+`, Text, nil},
			},
		}
	},
).SetAnalyser(func(text string) float32 {
	if text[:1] != "." {
		return 0
	}

	if text[:3] == `.\"` {
		return 1.0
	}

	if text[:4] == ".TH " {
		return 1.0
	}

	if groffAlphanumericRe.MatchString(text[1:3]) && unicode.IsSpace(rune(text[3])) {
		return 0.9
	}

	return 0
}))
