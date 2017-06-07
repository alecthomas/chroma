package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Makefile lexer.
var Makefile = Register(MustNewLexer(
	&Config{
		Name:      "Makefile",
		Aliases:   []string{"make", "makefile", "mf", "bsdmake"},
		Filenames: []string{"*.mak", "*.mk", "Makefile", "makefile", "Makefile.*", "GNUmakefile"},
		MimeTypes: []string{"text/x-makefile"},
	},
	Rules{
		"root": {
			{`^(?:[\t ]+.*\n|\n)+`, Using(Bash, nil), nil},
			{`\$[<@$+%?|*]`, Keyword, nil},
			{`\s+`, Text, nil},
			{`#.*?\n`, Comment, nil},
			{`(export)(\s+)`, ByGroups(Keyword, Text), Push("export")},
			{`export\s+`, Keyword, nil},
			{`([\w${}().-]+)(\s*)([!?:+]?=)([ \t]*)((?:.*\\\n)+|.*\n)`, ByGroups(NameVariable, Text, Operator, Text, Using(Bash, nil)), nil},
			{`(?s)"(\\\\|\\.|[^"\\])*"`, LiteralStringDouble, nil},
			{`(?s)'(\\\\|\\.|[^'\\])*'`, LiteralStringSingle, nil},
			{`([^\n:]+)(:+)([ \t]*)`, ByGroups(NameFunction, Operator, Text), Push("block-header")},
			{`\$\(`, Keyword, Push("expansion")},
		},
		"expansion": {
			{`[^$a-zA-Z_()]+`, Text, nil},
			{`[a-zA-Z_]+`, NameVariable, nil},
			{`\$`, Keyword, nil},
			{`\(`, Keyword, Push()},
			{`\)`, Keyword, Pop(1)},
		},
		"export": {
			{`[\w${}-]+`, NameVariable, nil},
			{`\n`, Text, Pop(1)},
			{`\s+`, Text, nil},
		},
		"block-header": {
			{`[,|]`, Punctuation, nil},
			{`#.*?\n`, Comment, Pop(1)},
			{`\\\n`, Text, nil},
			{`\$\(`, Keyword, Push("expansion")},
			{`[a-zA-Z_]+`, Name, nil},
			{`\n`, Text, Pop(1)},
			{`.`, Text, nil},
		},
	},
))
