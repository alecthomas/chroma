package i

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ini lexer.
var Ini = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "INI",
		Aliases:   []string{"ini", "cfg", "dosini"},
		Filenames: []string{"*.ini", "*.cfg", "*.inf", ".gitconfig", ".editorconfig"},
		MimeTypes: []string{"text/x-ini", "text/inf"},
	},
	iniRules,
).SetAnalyser(func(text string) float32 {
	npos := strings.Count(text, "\n")
	if npos < 3 {
		return 0
	}

	if text[0] == '[' && text[npos-1] == ']' {
		return 1.0
	}

	return 0
}))

func iniRules() Rules {
	return Rules{
		"root": {
			{`\s+`, Text, nil},
			{`[;#].*`, CommentSingle, nil},
			{`\[.*?\]$`, Keyword, nil},
			{`(.*?)([ \t]*)(=)([ \t]*)(.*(?:\n[ \t].+)*)`, ByGroups(NameAttribute, Text, Operator, Text, LiteralString), nil},
			{`(.+?)$`, NameAttribute, nil},
		},
	}
}
