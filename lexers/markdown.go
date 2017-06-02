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
				// ByGroups(Text, Keyword, Keyword, using(this, state='inline')), nil},
				ByGroups(Text, Keyword, Keyword, Text), nil},
			// bulleted lists
			{`^(\s*)([*-])(\s)(.+\n)`,
				// ByGroups(Text, Keyword, Text, using(this, state='inline')), nil},
				ByGroups(Text, Keyword, Text, Text), nil},
			// numbered lists
			{`^(\s*)([0-9]+\.)( .+\n)`,
				// ByGroups(Text, Keyword, using(this, state='inline')), nil},
				ByGroups(Text, Keyword, Text), nil},
			// quote
			{`^(\s*>\s)(.+\n)`, ByGroups(Keyword, GenericEmph), nil},
			// text block
			{"^(```\n)([\\w\\W]*?)(^```$)", ByGroups(String, Text, String), nil},
			// code block with language
			{"^(```)(\\w+)(\n)([\\w\\W]*?)(^```$)", EmitterFunc(HandleCodeblock), nil},
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

func HandleCodeblock(groups []string) []Token {
	out := []Token{
		{String, groups[1]},
		{String, groups[2]},
		{Text, groups[3]},
	}
	code := groups[4]
	lexer := Registry.Get(groups[2])
	tokens, err := lexer.Tokenise(code)
	if err == nil {
		out = append(out, tokens...)
	} else {
		out = append(out, Token{Error, code})
	}
	out = append(out, Token{String, groups[5]})
	return out

}
