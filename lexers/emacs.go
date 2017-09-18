package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Emacslisp lexer.
var Emacslisp = Register(MustNewLexer(
	&Config{
		Name:      "EmacsLisp",
		Aliases:   []string{"emacs", "elisp", "emacs-lisp"},
		Filenames: []string{"*.el"},
		MimeTypes: []string{"text/x-elisp", "application/x-elisp"},
	},
	Rules{
		"root": {
			Default(Push("body")),
		},
		"body": {
			{`\s+`, Text, nil},
			{`;.*$`, CommentSingle, nil},
			{`"`, LiteralString, Push("string")},
			{`\?([^\\]|\\.)`, LiteralStringChar, nil},
			{`:((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, NameBuiltin, nil},
			{`::((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, LiteralStringSymbol, nil},
			{`'((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, LiteralStringSymbol, nil},
			{`'`, Operator, nil},
			{"`", Operator, nil},
			{"[-+]?\\d+\\.?(?=[ \"()\\]\\'\\n,;`])", LiteralNumberInteger, nil},
			{"[-+]?\\d+/\\d+(?=[ \"()\\]\\'\\n,;`])", LiteralNumber, nil},
			{"[-+]?(\\d*\\.\\d+([defls][-+]?\\d+)?|\\d+(\\.\\d*)?[defls][-+]?\\d+)(?=[ \"()\\]\\'\\n,;`])", LiteralNumberFloat, nil},
			{`\[|\]`, Punctuation, nil},
			{`#:((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, LiteralStringSymbol, nil},
			{`#\^\^?`, Operator, nil},
			{`#\'`, NameFunction, nil},
			{`#[bB][+-]?[01]+(/[01]+)?`, LiteralNumberBin, nil},
			{`#[oO][+-]?[0-7]+(/[0-7]+)?`, LiteralNumberOct, nil},
			{`#[xX][+-]?[0-9a-fA-F]+(/[0-9a-fA-F]+)?`, LiteralNumberHex, nil},
			{`#\d+r[+-]?[0-9a-zA-Z]+(/[0-9a-zA-Z]+)?`, LiteralNumber, nil},
			{`#\d+=`, Operator, nil},
			{`#\d+#`, Operator, nil},
			{`(,@|,|\.|:)`, Operator, nil},
			{"(t|nil)(?=[ \"()\\]\\'\\n,;`])", NameConstant, nil},
			{`\*((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)\*`, NameVariableGlobal, nil},
			{`((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, NameVariable, nil},
			{`#\(`, Operator, Push("body")},
			{`\(`, Punctuation, Push("body")},
			{`\)`, Punctuation, Pop(1)},
		},
		"string": {
			{"[^\"\\\\`]+", LiteralString, nil},
			{"`((?:\\\\.|[\\w!$%&*+-/<=>?@^{}~|])(?:\\\\.|[\\w!$%&*+-/<=>?@^{}~|]|[#.:])*)\\'", LiteralStringSymbol, nil},
			{"`", LiteralString, nil},
			{`\\.`, LiteralString, nil},
			{`\\\n`, LiteralString, nil},
			{`"`, LiteralString, Pop(1)},
		},
	},
))
