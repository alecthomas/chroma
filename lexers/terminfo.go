package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Terminfo lexer.
var Terminfo = Register(MustNewLexer(
	&Config{
		Name:      "Terminfo",
		Aliases:   []string{"terminfo"},
		Filenames: []string{"terminfo", "terminfo.src"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {
			{`^#.*$`, Comment, nil},
			{`^[^\s#,|]+`, NameTag, Push("names")},
		},
		"names": {
			{`\n`, Text, Pop(1)},
			{`(,)([ \t]*)`, ByGroups(Punctuation, Text), Push("defs")},
			{`\|`, Punctuation, nil},
			{`[^,|]+`, NameAttribute, nil},
		},
		"defs": {
			{`\n[ \t]+`, Text, nil},
			{`\n`, Text, Pop(2)},
			{`(#)([0-9]+)`, ByGroups(Operator, LiteralNumber), nil},
			{`=`, Operator, Push("data")},
			{`(,)([ \t]*)`, ByGroups(Punctuation, Text), nil},
			{`[^\s,=#]+`, NameClass, nil},
		},
		"data": {
			{`\\[,\\]`, Literal, nil},
			{`(,)([ \t]*)`, ByGroups(Punctuation, Text), Pop(1)},
			{`[^\\,]+`, Literal, nil},
			{`.`, Literal, nil},
		},
	},
))
