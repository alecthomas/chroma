package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

var INI = Register(NewLexer(
	&Config{
		Name:      "INI",
		Aliases:   []string{"ini", "cfg", "dosini"},
		Filenames: []string{"*.ini", "*.cfg", "*.inf"},
		MimeTypes: []string{"text/x-ini", "text/inf"},
	},
	map[string][]Rule{
		"root": []Rule{
			{`\s+`, Whitespace, nil},
			{`;.*?$`, Comment, nil},
			{`\[.*?\]$`, Keyword, nil},
			{`(.*?)(\s*)(=)(\s*)(.*?)$`, ByGroups(Name, Whitespace, Operator, Whitespace, String), nil},
			// standalone option, supported by some INI parsers
			{`(.+?)$`, NameAttribute, nil},
		},
	},
))
