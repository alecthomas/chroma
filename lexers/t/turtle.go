package t

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var turtleAnalyserRe = regexp.MustCompile(`^\s*(@base|BASE|@prefix|PREFIX)`)

// Turtle lexer.
var Turtle = internal.Register(MustNewLazyLexer(
	&Config{
		Name:            "Turtle",
		Aliases:         []string{"turtle"},
		Filenames:       []string{"*.ttl"},
		MimeTypes:       []string{"text/turtle", "application/x-turtle"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	turtleRules,
).SetAnalyser(func(text string) float32 {
	// Turtle and Tera Term macro files share the same file extension
	// but each has a recognizable and distinct syntax.
	if turtleAnalyserRe.MatchString(text) {
		return 0.8
	}

	return 0
}))

func turtleRules() Rules {
	return Rules{
		"root": {
			{`\s+`, TextWhitespace, nil},
			{"(@base|BASE)(\\s+)(<[^<>\"{}|^`\\\\\\x00-\\x20]*>)(\\s*)(\\.?)", ByGroups(Keyword, TextWhitespace, NameVariable, TextWhitespace, Punctuation), nil},
			{"(@prefix|PREFIX)(\\s+)((?:[a-z][\\w-]*)?\\:)(\\s+)(<[^<>\"{}|^`\\\\\\x00-\\x20]*>)(\\s*)(\\.?)", ByGroups(Keyword, TextWhitespace, NameNamespace, TextWhitespace, NameVariable, TextWhitespace, Punctuation), nil},
			{`(?<=\s)a(?=\s)`, KeywordType, nil},
			{"(<[^<>\"{}|^`\\\\\\x00-\\x20]*>)", NameVariable, nil},
			{`((?:[a-z][\w-]*)?\:)([a-z][\w-]*)`, ByGroups(NameNamespace, NameTag), nil},
			{`#[^\n]+`, Comment, nil},
			{`\b(true|false)\b`, Literal, nil},
			{`[+\-]?\d*\.\d+`, LiteralNumberFloat, nil},
			{`[+\-]?\d*(:?\.\d+)?E[+\-]?\d+`, LiteralNumberFloat, nil},
			{`[+\-]?\d+`, LiteralNumberInteger, nil},
			{`[\[\](){}.;,:^]`, Punctuation, nil},
			{`"""`, LiteralString, Push("triple-double-quoted-string")},
			{`"`, LiteralString, Push("single-double-quoted-string")},
			{`'''`, LiteralString, Push("triple-single-quoted-string")},
			{`'`, LiteralString, Push("single-single-quoted-string")},
		},
		"triple-double-quoted-string": {
			{`"""`, LiteralString, Push("end-of-string")},
			{`[^\\]+`, LiteralString, nil},
			{`\\`, LiteralString, Push("string-escape")},
		},
		"single-double-quoted-string": {
			{`"`, LiteralString, Push("end-of-string")},
			{`[^"\\\n]+`, LiteralString, nil},
			{`\\`, LiteralString, Push("string-escape")},
		},
		"triple-single-quoted-string": {
			{`'''`, LiteralString, Push("end-of-string")},
			{`[^\\]+`, LiteralString, nil},
			{`\\`, LiteralString, Push("string-escape")},
		},
		"single-single-quoted-string": {
			{`'`, LiteralString, Push("end-of-string")},
			{`[^'\\\n]+`, LiteralString, nil},
			{`\\`, LiteralString, Push("string-escape")},
		},
		"string-escape": {
			{`.`, LiteralString, Pop(1)},
		},
		"end-of-string": {
			{`(@)([a-z]+(:?-[a-z0-9]+)*)`, ByGroups(Operator, GenericEmph, GenericEmph), Pop(2)},
			{"(\\^\\^)(<[^<>\"{}|^`\\\\\\x00-\\x20]*>)", ByGroups(Operator, GenericEmph), Pop(2)},
			{`(\^\^)((?:[a-z][\w-]*)?\:)([a-z][\w-]*)`, ByGroups(Operator, GenericEmph, GenericEmph), Pop(2)},
			Default(Pop(2)),
		},
	}
}
