package lexers

import (
	. "github.com/crowyy03/chroma/v2"
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
			{`#.*$`, CommentSingle, nil},
			{`"([^"\\]|\\.)*"`, LiteralString, nil},
			{`[0-9]+(\.[0-9]+)?([eE][+-]?[0-9]+)?`, LiteralNumber, nil},
			{`0x[0-9a-fA-F]+`, LiteralNumberHex, nil},
			{`[+\-*/%<>=!&|^~@]`, Operator, nil},
			{`[\[\](){}.,;:]`, Punctuation, nil},
			{`\b(use|toArray|toString|joinPath|set|get|new|cast|neg|rshift|and|cat|times|dup|between|within|when|assert|pfunc|overload|same|isCombined|meetsAll|loadString|saveFile|printList|failProc|raiseStaticError|lexicalError|dynamic|struct|fieldCount|toSpan2|toSpanStatic2|toStringView|getCodePointAndSize|splitString|assembleString|uif|while|if|exportFunction|ensure|toCommandLine2|loadString|saveFile|printList|meetsAll|pfunc|overload|same|isCombined|TRUE|FALSE)\b`, Keyword, nil},
			{`\b(Array|String|Int32|Int64|Nat8|Nat32|Natx|Real64|HashTable|Variant|CommandLine|Function|Process|Text|Cref|Ref|Cond|StringView)\b`, KeywordType, nil},
			{`\b(LF|CR|TAB)\b`, LiteralStringEscape, nil},
			{`[a-zA-Z_][a-zA-Z0-9_]*:`, NameFunction, nil},
			{`[a-zA-Z_][a-zA-Z0-9_]*`, Name, nil},
			{`\s+`, TextWhitespace, nil},
		},
	}
}
