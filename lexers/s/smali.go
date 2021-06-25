package s

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	smaliAnalyserClassRe         = regexp.MustCompile(`(?m)^\s*\.class\s`)
	smaliAnalyserClassKeywordsRe = regexp.MustCompile(
		`(?m)\b((check-cast|instance-of|throw-verification-error` +
			`)\b|(-to|add|[ais]get|[ais]put|and|cmpl|const|div|` +
			`if|invoke|move|mul|neg|not|or|rem|return|rsub|shl` +
			`|shr|sub|ushr)[-/])|{|}`)
	smaliAnalyserKeywordsRe = regexp.MustCompile(
		`(?m)(\.(catchall|epilogue|restart local|prologue)|` +
			`\b(array-data|class-change-error|declared-synchronized|` +
			`(field|inline|vtable)@0x[0-9a-fA-F]|generic-error|` +
			`illegal-class-access|illegal-field-access|` +
			`illegal-method-access|instantiation-error|no-error|` +
			`no-such-class|no-such-field|no-such-method|` +
			`packed-switch|sparse-switch))\b`)
)

// Smali lexer.
var Smali = internal.Register(MustNewLexer(
	&Config{
		Name:      "Smali",
		Aliases:   []string{"smali"},
		Filenames: []string{"*.smali"},
		MimeTypes: []string{"text/smali"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	var result float32

	if smaliAnalyserClassRe.MatchString(text) {
		result += 0.5

		if smaliAnalyserClassKeywordsRe.MatchString(text) {
			result += 0.3
		}
	}

	if smaliAnalyserKeywordsRe.MatchString(text) {
		result += 0.6
	}

	return result
}))
