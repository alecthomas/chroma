package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Markdown lexer.
var Markdown = Register(NewLexer(
	&Config{
		Name:      "markdown",
		Aliases:   []string{"md"},
		Filenames: []string{"*.md"},
		MimeTypes: []string{"text/x-markdown"},
	},
	map[string][]Rule{
		"root": []Rule{
			// heading with pound prefix
			{`^(#)([^#].+\n)`, ByGroups(GenericHeading, Text), nil},
			{`^(#{2,6})(.+\n)`, ByGroups(GenericSubheading, Text), nil},
			// task list
			{`^(\s*)([*-] )(\[[ xX]\])( .+\n)`,
				ByGroups(Text, Keyword, Keyword, Text), nil},
			// bulleted lists
			{`^(\s*)([*-])(\s)(.+\n)`,
				ByGroups(Text, Keyword, Text, Text), nil},
			// numbered lists
			{`^(\s*)([0-9]+\.)( .+\n)`,
				ByGroups(Text, Keyword, Text), nil},
			// quote
			{`^(\s*>\s)(.+\n)`, ByGroups(Keyword, GenericEmph), nil},
			// text block
			{"^(```\n)([\\w\\W]*?)(^```$)", ByGroups(String, Text, String), nil},
			// code block with language
			{"^(```)(\\w+)(\n)([\\w\\W]*?)(^```$)", EmitterFunc(handleCodeblock), nil},
			Include(`inline`),
		},
		`inline`: []Rule{
			// escape
			{`\\.`, Text, nil},
			// italics
			{`(\s)([*_][^*_]+[*_])(\W|\n)`, ByGroups(Text, GenericEmph, Text), nil},
			// bold
			// warning: the following rule eats internal tags. eg. **foo _bar_ baz** bar is not italics
			{`(\s)(\*\*.*\*\*)`, ByGroups(Text, GenericStrong), nil},
			// strikethrough
			{`(\s)(~~[^~]+~~)`, ByGroups(Text, GenericDeleted), nil},
			// inline code
			{"`[^`]+`", StringBacktick, nil},
			// mentions and topics (twitter and github stuff)
			{`[@#][\w/:]+`, NameEntity, nil},
			// (image?) links eg: ![Image of Yaktocat](https://octodex.github.com/images/yaktocat.png)
			{`(!?\[)([^]]+)(\])(\()([^)]+)(\))`, ByGroups(Text, NameTag, Text, Text, NameAttribute, Text), nil},

			// general text, must come last!
			{`[^\\\s]+`, Text, nil},
			{`.`, Text, nil},
		},
	},
))

func handleCodeblock(groups []string, out func(Token)) {
	out(Token{String, groups[1]})
	out(Token{String, groups[2]})
	out(Token{Text, groups[3]})
	code := groups[4]
	lexer := Registry.Get(groups[2])
	lexer.Tokenise(code, out)
	out(Token{String, groups[5]})
}
