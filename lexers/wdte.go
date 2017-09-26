package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// WDTE lexer.
var WDTE = Register(MustNewLexer(
	&Config{
		Name:      "WDTE",
		Filenames: []string{"*.wdte"},
	},
	Rules{
		"root": {
			{`\n`, Text, nil},
			{`\s+`, Text, nil},
			{`\\\n`, Text, nil},
			{`#(.*?)\n`, CommentSingle, nil},
			{`-?[0-9]+`, LiteralNumberInteger, nil},
			{`-?[0-9]*\.[0-9]+`, LiteralNumberFloat, nil},
			{`"[^"]*"`, LiteralString, nil},
			{`'[^']*'`, LiteralString, nil},
			{Words(``, `\b`, `switch`, `default`, `memo`), KeywordReserved, nil},
			{`{|}|;|->|=>|\(|\)|\[|\]|\.`, Operator, nil},
			{`[^{};()[\].\s]+`, NameVariable, nil},
		},
	},
))
