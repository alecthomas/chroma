package lexers

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint: golint
)

// Go lexer.
var Go = Register(MustNewLexer(
	&Config{
		Name:      "Go",
		Filenames: []string{"*.go"},
		Aliases:   []string{"go", "golang"},
		MimeTypes: []string{"text/x-gosrc"},
	},
	// TODO: Convert this Lexer to use text/scanner
	Rules{
		`root`: []Rule{
			{`\n`, Text, nil},
			{`\s+`, Text, nil},
			{`\\\n`, Text, nil}, // line continuations
			{`//(.*?)\n`, CommentSingle, nil},
			{`/(\\\n)?[*](.|\n)*?[*](\\\n)?/`, CommentMultiline, nil},
			{`(import|package)\b`, KeywordNamespace, nil},
			{`(var|func|struct|map|chan|type|interface|const)\b`,
				KeywordDeclaration, nil},
			{Words(`break`, `default`, `select`, `case`, `defer`, `go`,
				`else`, `goto`, `switch`, `fallthrough`, `if`, `range`,
				`continue`, `for`, `return`), Keyword, nil},
			{`(true|false|iota|nil)\b`, KeywordConstant, nil},
			{Words(`uint`, `uint8`, `uint16`, `uint32`, `uint64`,
				`int`, `int8`, `int16`, `int32`, `int64`,
				`float`, `float32`, `float64`,
				`complex64`, `complex128`, `byte`, `rune`,
				`string`, `bool`, `erro`, `uintpt`,
				`print`, `println`, `panic`, `recove`, `close`, `complex`,
				`real`, `imag`, `len`, `cap`, `append`, `copy`, `delete`,
				`new`, `make`),
				KeywordType, nil},
			// imaginary_lit
			{`\d+i`, LiteralNumber, nil},
			{`\d+\.\d*([Ee][-+]\d+)?i`, LiteralNumber, nil},
			{`\.\d+([Ee][-+]\d+)?i`, LiteralNumber, nil},
			{`\d+[Ee][-+]\d+i`, LiteralNumber, nil},
			// float_lit
			{`\d+(\.\d+[eE][+\-]?\d+|\.\d*|[eE][+\-]?\d+)`, LiteralNumberFloat, nil},
			{`\.\d+([eE][+\-]?\d+)?`, LiteralNumberFloat, nil},
			// int_lit
			// -- octal_lit
			{`0[0-7]+`, LiteralNumberOct, nil},
			// -- hex_lit
			{`0[xX][0-9a-fA-F]+`, LiteralNumberHex, nil},
			// -- decimal_lit
			{`(0|[1-9][0-9]*)`, LiteralNumberInteger, nil},
			// char_lit
			{`'(\\['"\\abfnrtv]|\\x[0-9a-fA-F]{2}|\\[0-7]{1,3}|\\u[0-9a-fA-F]{4}|\\U[0-9a-fA-F]{8}|[^\\])'`, LiteralStringChar, nil},
			// StringLiteral
			// -- raw_string_lit
			{"`[^`]*`", String, nil},
			// -- interpreted_string_lit
			{`"(\\\\|\\"|[^"])*"`, String, nil},
			// Tokens
			{`(<<=|>>=|<<|>>|<=|>=|&\^=|&\^|\+=|-=|\*=|/=|%=|&=|\|=|&&|\|\||<-|\+\+|--|==|!=|:=|\.\.\.|[+\-*/%&])`, Operator, nil},
			{`[|^<>=!()\[\]{}.,;:]`, Punctuation, nil},
			// identifier
			{`[^\W\d]\w*`, NameOther, nil},
		},
	},
).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, "fmt.") && strings.Contains(text, "package ") {
		return 0.5
	}
	if strings.Contains(text, "package ") {
		return 0.1
	}
	return 0.0
}))
