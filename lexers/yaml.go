package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

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
