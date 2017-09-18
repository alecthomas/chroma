package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Termcap lexer.
var Termcap = Register(MustNewLexer(
	&Config{
		Name:      "Termcap",
		Aliases:   []string{"termcap"},
		Filenames: []string{"termcap", "termcap.src"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {
			{`^#.*$`, Comment, nil},
			{`^[^\s#:|]+`, NameTag, Push("names")},
		},
		"names": {
			{`\n`, Text, Pop(1)},
			{`:`, Punctuation, Push("defs")},
			{`\|`, Punctuation, nil},
			{`[^:|]+`, NameAttribute, nil},
		},
		"defs": {
			{`\\\n[ \t]*`, Text, nil},
			{`\n[ \t]*`, Text, Pop(2)},
			{`(#)([0-9]+)`, ByGroups(Operator, LiteralNumber), nil},
			{`=`, Operator, Push("data")},
			{`:`, Punctuation, nil},
			{`[^\s:=#]+`, NameClass, nil},
		},
		"data": {
			{`\\072`, Literal, nil},
			{`:`, Punctuation, Pop(1)},
			{`[^:\\]+`, Literal, nil},
			{`.`, Literal, nil},
		},
	},
))
