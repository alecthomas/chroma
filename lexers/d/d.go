package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// D lexer. See https://dlang.org/spec/lex.html
var D = internal.Register(MustNewLexer(
	&Config{
		Name:      "D",
		Aliases:   []string{"d"},
		Filenames: []string{"*.d", "*.di"},
		MimeTypes: []string{"text/x-d"},
		EnsureNL:  true,
	},
	Rules{
		"root": {
			{`[^\S\n]+`, Text, nil},
			// https://dlang.org/spec/lex.html#comment
			{`//.*?\n`, CommentSingle, nil},
			{`/\*.*?\*/`, CommentMultiline, nil},
			{`/\+.*?\+/`, CommentMultiline, nil},
			// https://dlang.org/spec/lex.html#keywords
			{`(asm|assert|body|break|case|catch|continue|default|do|else|finally|for|if|goto|instanceof|new|return|switch|this|throw|try|while)\b`, Keyword, nil},
			{`((?:(?:[^\W\d]|\$)[\w.\[\]$<>]*\s+)+?)((?:[^\W\d]|\$)[\w$]*)(\s*)(\()`, ByGroups(UsingSelf("root"), NameFunction, Text, Operator), nil},
			{`@[^\W\d][\w.]*`, NameDecorator, nil},
			{`(abstract|auto|alias|align|const|enum|extends|final|package|private|protected|public|static|strictfp|super|synchronized|throws|transient|volatile)\b`, KeywordDeclaration, nil},
			{`(bool|byte|char|double|float|int|long|short|void)\b`, KeywordType, nil},
			{`(module)(\s+)`, ByGroups(KeywordNamespace, Text), Push("import")},
			{`(true|false|null)\b`, KeywordConstant, nil},
			{`(class|interface)(\s+)`, ByGroups(KeywordDeclaration, Text), Push("class")},
			{`(import(?:\s+static)?)(\s+)`, ByGroups(KeywordNamespace, Text), Push("import")},
			{`"(\\\\|\\"|[^"])*"`, LiteralString, nil},
			{`'\\.'|'[^\\]'|'\\u[0-9a-fA-F]{4}'`, LiteralStringChar, nil},
			{`(\.)((?:[^\W\d]|\$)[\w$]*)`, ByGroups(Operator, NameAttribute), nil},
			{`^\s*([^\W\d]|\$)[\w$]*:`, NameLabel, nil},
			{`([0-9][0-9_]*\.([0-9][0-9_]*)?|\.[0-9][0-9_]*)([eE][+\-]?[0-9][0-9_]*)?[fFdD]?|[0-9][eE][+\-]?[0-9][0-9_]*[fFdD]?|[0-9]([eE][+\-]?[0-9][0-9_]*)?[fFdD]|0[xX]([0-9a-fA-F][0-9a-fA-F_]*\.?|([0-9a-fA-F][0-9a-fA-F_]*)?\.[0-9a-fA-F][0-9a-fA-F_]*)[pP][+\-]?[0-9][0-9_]*[fFdD]?`, LiteralNumberFloat, nil},
			{`0[xX][0-9a-fA-F][0-9a-fA-F_]*[lL]?`, LiteralNumberHex, nil},
			{`0[bB][01][01_]*[lL]?`, LiteralNumberBin, nil},
			{`0[0-7_]+[lL]?`, LiteralNumberOct, nil},
			{`0|[1-9][0-9_]*[lL]?`, LiteralNumberInteger, nil},
			{`[~^*!%&\[\](){}<>|+=:;,./?-]`, Operator, nil},
			{`([^\W\d]|\$)[\w$]*`, Name, nil},
			{`\n`, Text, nil},
		},
		"class": {
			{`([^\W\d]|\$)[\w$]*`, NameClass, Pop(1)},
		},
		"import": {
			{`[\w.]+\*?`, NameNamespace, Pop(1)},
		},
	},
))
