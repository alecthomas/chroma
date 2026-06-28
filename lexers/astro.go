package lexers

import (
	. "github.com/alecthomas/chroma/v3" // nolint
)

// Astro lexer.
var Astro = Register(DelegatingLexer(HTML, MustNewLexer(
	&Config{
		Name:            "Astro",
		Aliases:         []string{"astro"},
		Filenames:       []string{"*.astro"},
		MimeTypes:       []string{"text/x-astro"},
		CaseInsensitive: true,
		DotAll:          true,
	},
	astroRules,
)))

func astroRules() Rules {
	return Rules{
		"root": {
			{
				`(\A---[ \t]*\r?\n)(.*?)(^---[ \t]*(?:\r?\n|$))`,
				ByGroups(CommentPreproc, Using("TypeScript"), CommentPreproc),
				nil,
			},
			// Let HTML handle comments, including comments containing script and style tags.
			{`<!--`, Other, Push("comment")},
			{
				`(<\s*script\b(?=[^>]*(?:lang|type)\s*=\s*['"]?(?:json|application/(?:ld\+)?json|importmap|speculationrules)(?:['"]|[\s>]))[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*script\s*>)`,
				ByGroups(Other, Using("JSON"), Other),
				nil,
			},
			{
				`(<\s*script\b(?=[^>]*(?:lang|type)\s*=\s*['"]?(?:js|javascript|text/javascript|application/(?:x-)?javascript|text/ecmascript|application/ecmascript|module)(?:['"]|[\s>]))[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*script\s*>)`,
				ByGroups(Other, Using("JavaScript"), Other),
				nil,
			},
			{
				`(<\s*script\b(?=[^>]*(?:lang|type)\s*=\s*['"]?(?:ts|typescript|text/typescript|application/typescript)(?:['"]|[\s>]))[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*script\s*>)`,
				ByGroups(Other, Using("TypeScript"), Other),
				nil,
			},
			{
				`(<\s*script\b[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*script\s*>)`,
				ByGroups(Other, Using("TypeScript"), Other),
				nil,
			},
			{
				`(<\s*style\b(?=[^>]*lang\s*=\s*['"]?(?:scss|source\.scss|source\.css\.scss)(?:['"]|[\s>]))[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*style\s*>)`,
				ByGroups(Other, Using("SCSS"), Other),
				nil,
			},
			{
				`(<\s*style\b(?=[^>]*lang\s*=\s*['"]?(?:sass|source\.sass)(?:['"]|[\s>]))[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*style\s*>)`,
				ByGroups(Other, Using("Sass"), Other),
				nil,
			},
			{
				`(<\s*style\b(?=[^>]*lang\s*=\s*['"]?(?:stylus|source\.stylus)(?:['"]|[\s>]))[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*style\s*>)`,
				ByGroups(Other, Using("Stylus"), Other),
				nil,
			},
			{
				`(<\s*style\b[^>]*>)` +
					`(.*?)` +
					`(<\s*/\s*style\s*>)`,
				ByGroups(Other, Using("CSS"), Other),
				nil,
			},
			{
				`(\s*)(=)(\s*)({)`,
				ByGroups(Text, Operator, Text, Punctuation),
				Push("expression"),
			},
			{`{`, Punctuation, Push("expression")},
			{`.+?`, Other, nil},
		},
		"comment": {
			{`-->`, Other, Pop(1)},
			{`.+?`, Other, nil},
		},
		"expression": {
			{`}`, Punctuation, Pop(1)},
			// Let TypeScript handle strings and the curly braces inside them.
			{`(?<!(?<!\\)\\)(['"` + "`])" + `.*?(?<!(?<!\\)\\)\1`, Using("TypeScript"), nil},
			{"{", Punctuation, Push("expression")},
			{`[^{}]+`, Using("TypeScript"), nil},
		},
	}
}
