package b

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Brainfuck lexer.
var Brainfuck = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Brainfuck",
		Aliases:   []string{"brainfuck", "bf"},
		Filenames: []string{"*.bf", "*.b"},
		MimeTypes: []string{"application/x-brainfuck"},
	},
	brainfuckRules,
).SetAnalyser(func(text string) float32 {
	// it's safe to assume that a program which mostly consists of + -
	// and < > is brainfuck.
	var plusMinusCount float64
	var greaterLessCount float64

	rangeToCheck := len(text)

	if rangeToCheck > 256 {
		rangeToCheck = 256
	}

	for _, c := range text[:rangeToCheck] {
		if c == '+' || c == '-' {
			plusMinusCount++
		}
		if c == '<' || c == '>' {
			greaterLessCount++
		}
	}

	if plusMinusCount > (0.25 * float64(rangeToCheck)) {
		return 1.0
	}

	if greaterLessCount > (0.25 * float64(rangeToCheck)) {
		return 1.0
	}

	if strings.Contains(text, "[-]") {
		return 0.5
	}

	return 0
}))

func brainfuckRules() Rules {
	return Rules{
		"common": {
			{`[.,]+`, NameTag, nil},
			{`[+-]+`, NameBuiltin, nil},
			{`[<>]+`, NameVariable, nil},
			{`[^.,+\-<>\[\]]+`, Comment, nil},
		},
		"root": {
			{`\[`, Keyword, Push("loop")},
			{`\]`, Error, nil},
			Include("common"),
		},
		"loop": {
			{`\[`, Keyword, Push()},
			{`\]`, Keyword, Pop(1)},
			Include("common"),
		},
	}
}
