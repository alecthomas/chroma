package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Prism.languages.yaml = {
// 	'scalar': {
// 		pattern: /([\-:]\s*(![^\s]+)?[ \t]*[|>])[ \t]*(?:((?:\r?\n|\r)[ \t]+)[^\r\n]+(?:\3[^\r\n]+)*)/,
// 		lookbehind: true,
// 		alias: 'string'
// 	},
// 	'comment': /#.*/,
// 	'key': {
// 		pattern: /(\s*(?:^|[:\-,[{\r\n?])[ \t]*(![^\s]+)?[ \t]*)[^\r\n{[\]},#\s]+?(?=\s*:\s)/,
// 		lookbehind: true,
// 		alias: 'atrule'
// 	},
// 	'directive': {
// 		pattern: /(^[ \t]*)%.+/m,
// 		lookbehind: true,
// 		alias: 'important'
// 	},
// 	'datetime': {
// 		pattern: /([:\-,[{]\s*(![^\s]+)?[ \t]*)(\d{4}-\d\d?-\d\d?([tT]|[ \t]+)\d\d?:\d{2}:\d{2}(\.\d*)?[ \t]*(Z|[-+]\d\d?(:\d{2})?)?|\d{4}-\d{2}-\d{2}|\d\d?:\d{2}(:\d{2}(\.\d*)?)?)(?=[ \t]*($|,|]|}))/m,
// 		lookbehind: true,
// 		alias: 'number'
// 	},
// 	'boolean': {
// 		pattern: /([:\-,[{]\s*(![^\s]+)?[ \t]*)(true|false)[ \t]*(?=$|,|]|})/im,
// 		lookbehind: true,
// 		alias: 'important'
// 	},
// 	'null': {
// 		pattern: /([:\-,[{]\s*(![^\s]+)?[ \t]*)(null|~)[ \t]*(?=$|,|]|})/im,
// 		lookbehind: true,
// 		alias: 'important'
// 	},
// 	'string': {
// 		pattern: /([:\-,[{]\s*(![^\s]+)?[ \t]*)("(?:[^"\\]|\\.)*"|'(?:[^'\\]|\\.)*')(?=[ \t]*($|,|]|}))/m,
// 		lookbehind: true,
// 		greedy: true
// 	},
// 	'number': {
// 		pattern: /([:\-,[{]\s*(![^\s]+)?[ \t]*)[+\-]?(0x[\da-f]+|0o[0-7]+|(\d+\.?\d*|\.?\d+)(e[\+\-]?\d+)?|\.inf|\.nan)[ \t]*(?=$|,|]|})/im,
// 		lookbehind: true
// 	},
// 	'tag': /![^\s]+/,
// 	'important': /[&*][\w]+/,
// 	'punctuation': /---|[:[\]{}\-,|>?]|\.\.\./

var YAML = Register(MustNewLexer(
	&Config{
		Name:      "YAML",
		Aliases:   []string{"yaml"},
		Filenames: []string{"*.yaml", "*.yml"},
		MimeTypes: []string{"text/x-yaml"},
	},
	Rules{
		"root": {
			{`\s+`, Whitespace, nil},
			{`#.*`, Comment, nil},
			{`!![^\s]+`, CommentPreproc, nil},
			{`&[^\s]+`, CommentPreproc, nil},
			{`\*[^\s]+`, CommentPreproc, nil},
			{Words(``, `\b`, "true", "false", "null"), KeywordConstant, nil},
			{`"(?:\\.|[^"])+"`, StringDouble, nil},
			{`\d\d\d\d-\d\d-\d\d([T ]\d\d:\d\d:\d\d(\.\d+)?(Z|\s+[-+]\d+)?)?`, LiteralDate, nil},
			{`([>|])(\s+)((?:(?:.*?$)(?:[\n\r]*?\2)?)*)`, ByGroups(StringDoc, StringDoc, StringDoc), nil},
			{`[+\-]?(0x[\da-f]+|0o[0-7]+|(\d+\.?\d*|\.?\d+)(e[\+\-]?\d+)?|\.inf|\.nan)\b`, Number, nil},
			{`[?:,\[\]]`, Punctuation, nil},
			{`.`, Text, nil},
		},
	},
))
