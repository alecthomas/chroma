package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Ini lexer.
var Ini = Register(MustNewLexer(
	&Config{
		Name:      "INI",
		Aliases:   []string{"ini", "cfg", "dosini"},
		Filenames: []string{"*.ini", "*.cfg", "*.inf"},
		MimeTypes: []string{"text/x-ini", "text/inf"},
	},
	Rules{
		"root": {
			{`\s+`, Text, nil},
			{`[;#].*`, CommentSingle, nil},
			{`\[.*?\]$`, Keyword, nil},
			{`(.*?)([ \t]*)(=)([ \t]*)(.*(?:\n[ \t].+)*)`, ByGroups(NameAttribute, Text, Operator, Text, LiteralString), nil},
			{`(.+?)$`, NameAttribute, nil},
		},
	},
))
