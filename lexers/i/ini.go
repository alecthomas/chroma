package i

import (
	. "github.com/colin3dmax/chroma" // nolint
	"github.com/colin3dmax/chroma/lexers/internal"
)

// Ini lexer.
var Ini = internal.Register(MustNewLexer(
	&Config{
		Name:      "INI",
		Aliases:   []string{"ini", "cfg", "dosini"},
		Filenames: []string{"*.ini", "*.cfg", "*.inf", ".gitconfig"},
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
