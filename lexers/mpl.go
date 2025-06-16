package lexers

import (
	. "github.com/alecthomas/chroma/v2"
)

// MPL lexer.
var MPL = Register(MustNewLexer(
	&Config{
		Name:      "MPL",
		Aliases:   []string{"mpl"},
		Filenames: []string{"*.mpl"},
		MimeTypes: []string{"text/x-mpl"},
	},
	mplRules,
))

func mplRules() Rules {
	return Rules{
		"root": {
			{`#.*`, CommentSingle, nil},
			{`\s+`, Text, nil},
			{`"(\\.|[^"])*"`, LiteralString, nil},
			{`'(\\.|[^'])*'`, LiteralString, nil},
			{`0x[0-9a-fA-F]+`, LiteralNumberHex, nil},
			{`\d+\.\d+([eE][+-]?\d+)?`, LiteralNumberFloat, nil},
			{`\d+`, LiteralNumberInteger, nil},
			{Words(``, `\b`,
				"if", "when", "while", "for", "use", "importFunction", "exportFunction",
				"return", "break", "continue", "throw", "try", "catch", "finally",
				"module", "package", "import", "export", "public", "private", "protected",
				"static", "dynamic", "const", "var", "let", "type", "interface", "class",
				"struct", "enum", "union", "typedef", "namespace", "using", "as",
			), Keyword, nil},
			{Words(``, `\b`,
				"Int8", "Int16", "Int32", "Int64", "UInt8", "UInt16", "UInt32", "UInt64",
				"Nat8", "Nat16", "Nat32", "Nat64", "Real32", "Real64", "Bool", "Char",
				"String", "Array", "List", "Map", "Set", "Queue", "Stack", "Vector",
				"Matrix", "Complex", "Rational", "Decimal", "Date", "Time", "DateTime",
				"Duration", "UUID", "Path", "File", "Directory", "Stream", "Buffer",
				"Socket", "Thread", "Process", "Mutex", "Semaphore", "Condition",
				"Event", "Future", "Promise", "Result", "Option", "Either", "Try",
				"Ref", "Ptr", "SmartPtr", "UniquePtr", "SharedPtr", "WeakPtr",
				"Owner", "Borrower", "View", "Span", "Range", "Iterator", "Generator",
				"Function", "Lambda", "Closure", "Coroutine", "Async", "Await",
				"XMLValue", "XMLDocument", "XMLElement", "XMLAttribute", "XMLParserResult",
				"XMLParserErrorInfo", "XMLParserPosition", "XMLVALUE_CHARDATA", "XMLVALUE_ELEMENT",
				"TRUE", "FALSE", "NULL", "NONE", "SOME", "OK", "ERROR",
			), KeywordType, nil},
			{`[\[\]\{\}\(\)\+\-\*/=<>!&|~^%:,;\.]`, Operator, nil},
			{`[A-Za-z_][A-Za-z0-9_\.]*`, Name, nil},
		},
	}
}
