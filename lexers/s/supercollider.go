package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Supercollider lexer.
var Supercollider = internal.Register(MustNewLexer(
	&Config{
		Name:      "SuperCollider",
		Aliases:   []string{"sc", "supercollider"},
		Filenames: []string{"*.sc", "*.scd"},
		MimeTypes: []string{"application/supercollider", "text/supercollider"},
		DotAll:    true,
	},
	Rules{
		"commentsandwhitespace": {
			{`\s+`, Text, nil},
			{`<!--`, Comment, nil},
			{`//.*?\n`, CommentSingle, nil},
			{`/\*.*?\*/`, CommentMultiline, nil},
		},
		"slashstartsregex": {
			Include("commentsandwhitespace"),
			{`/(\\.|[^[/\\\n]|\[(\\.|[^\]\\\n])*])+/([gim]+\b|\B)`, LiteralStringRegex, Pop(1)},
			{`(?=/)`, Text, Push("#pop", "badregex")},
			Default(Pop(1)),
		},
		"badregex": {
			{`\n`, Text, Pop(1)},
		},
		"root": {
			{`^(?=\s|/|<!--)`, Text, Push("slashstartsregex")},
			Include("commentsandwhitespace"),
			{`\+\+|--|~|&&|\?|:|\|\||\\(?=\n)|(<<|>>>?|==?|!=?|[-<>+*%&|^/])=?`, Operator, Push("slashstartsregex")},
			{`[{(\[;,]`, Punctuation, Push("slashstartsregex")},
			{`[})\].]`, Punctuation, nil},
			{Words(``, `\b`, `for`, `in`, `while`, `do`, `break`, `return`, `continue`, `switch`, `case`, `default`, `if`, `else`, `throw`, `try`, `catch`, `finally`, `new`, `delete`, `typeof`, `instanceof`, `void`), Keyword, Push("slashstartsregex")},
			{Words(``, `\b`, `var`, `let`, `with`, `function`, `arg`), KeywordDeclaration, Push("slashstartsregex")},
			{Words(``, `\b`, `(abstract`, `boolean`, `byte`, `char`, `class`, `const`, `debugger`, `double`, `enum`, `export`, `extends`, `final`, `float`, `goto`, `implements`, `import`, `int`, `interface`, `long`, `native`, `package`, `private`, `protected`, `public`, `short`, `static`, `super`, `synchronized`, `throws`, `transient`, `volatile`), KeywordReserved, nil},
			{Words(``, `\b`, `true`, `false`, `nil`, `inf`), KeywordConstant, nil},
			{Words(``, `\b`, `Array`, `Boolean`, `Date`, `Error`, `Function`, `Number`, `Object`, `Packages`, `RegExp`, `String`, `isFinite`, `isNaN`, `parseFloat`, `parseInt`, `super`, `thisFunctionDef`, `thisFunction`, `thisMethod`, `thisProcess`, `thisThread`, `this`), NameBuiltin, nil},
			{`[$a-zA-Z_]\w*`, NameOther, nil},
			{`\\?[$a-zA-Z_]\w*`, LiteralStringSymbol, nil},
			{`[0-9][0-9]*\.[0-9]+([eE][0-9]+)?[fd]?`, LiteralNumberFloat, nil},
			{`0x[0-9a-fA-F]+`, LiteralNumberHex, nil},
			{`[0-9]+`, LiteralNumberInteger, nil},
			{`"(\\\\|\\"|[^"])*"`, LiteralStringDouble, nil},
			{`'(\\\\|\\'|[^'])*'`, LiteralStringSingle, nil},
		},
	},
))
