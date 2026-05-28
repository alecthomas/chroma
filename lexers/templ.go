package lexers

import (
	"strings"

	. "github.com/alecthomas/chroma/v2" // nolint
)

// Templ lexer.
var Templ = Register(DelegatingLexer(HTML, MustNewLexer(
	&Config{
		Name:      "Templ",
		Aliases:   []string{"templ"},
		Filenames: []string{"*.templ"},
		MimeTypes: []string{"text/x-templ"},
		DotAll:    true,
	},
	templRules,
)).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, "templ ") && strings.Contains(text, "{") {
		return 0.7
	}
	if strings.Contains(text, "package ") && strings.Contains(text, "templ ") {
		return 0.5
	}
	return 0.0
}))

func templRules() Rules {
	goKeywords := Words(``, `\b`,
		"break", "case", "chan", "const", "continue", "default", "defer", "else",
		"fallthrough", "for", "func", "go", "goto", "if", "import", "interface",
		"map", "package", "range", "return", "select", "struct", "switch", "type", "var",
	)
	goTypes := Words(``, `\b`,
		"any", "bool", "byte", "comparable", "complex64", "complex128", "error",
		"float32", "float64", "int", "int8", "int16", "int32", "int64", "rune",
		"string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	)

	return Rules{
		"root": {
			{`\n`, TextWhitespace, nil},
			{`\s+`, TextWhitespace, nil},
			{`//[^\n\r]*`, CommentSingle, nil},
			{`/\*(?:.|\n)*?\*/`, CommentMultiline, nil},
			{`"(?:\\.|[^"\\])*"`, LiteralStringDouble, nil},
			{"`[^`]*`", LiteralStringBacktick, nil},
			{`'(?:\\.|[^'\\])'`, LiteralStringChar, nil},
			{`</?[A-Za-z][^>{}]*>`, Other, nil},
			{`<![^>{}]*>`, Other, nil},
			{`(@)([A-Za-z_]\w*)`, ByGroups(Punctuation, NameFunction), nil},
			{`\b(templ)(\s+)([A-Za-z_]\w*)`, ByGroups(KeywordDeclaration, TextWhitespace, NameFunction), nil},
			{goKeywords, Keyword, nil},
			{goTypes, KeywordType, nil},
			{`\b(true|false|nil|iota)\b`, KeywordConstant, nil},
			{`\b\d+(?:\.\d+)?(?:[eE][+\-]?\d+)?\b`, LiteralNumber, nil},
			{`[A-Z][A-Za-z0-9_]*`, NameClass, nil},
			{`[a-zA-Z_]\w*(?=\s*\()`, NameFunction, nil},
			{`[a-zA-Z_]\w*`, NameOther, nil},
			{`:=|==|!=|<=|>=|&&|\|\||\+\+|--|[-+*/%&|^!<>]=?`, Operator, nil},
			{`[{}()[\].,;:]`, Punctuation, nil},
			{`[^\s]+`, Other, nil},
		},
	}
}
