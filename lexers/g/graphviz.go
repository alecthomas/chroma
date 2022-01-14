package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Graphviz lexer.
var Graphviz = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Graphviz",
		Aliases:   []string{"graphviz", "dot"},
		Filenames: []string{"*.gv", "*.dot"},
		MimeTypes: []string{"text/x-graphviz", "text/vnd.graphviz"},
	},
	func() Rules {
		return Rules{
			"root": {
				{`\s+`, TextWhitespace, nil},
				{`(#|//).*?$`, CommentSingle, nil},
				{`/(\\\n)?[*](.|\n)*?[*](\\\n)?/`, CommentMultiline, nil},
				{`(?i)(node|edge|graph|digraph|subgraph|strict)\b`, Keyword, nil},
				{`--|->`, Operator, nil},
				{`[{}[\]:;,]`, Punctuation, nil},
				{`(\b\D\w*)(\s*)(=)(\s*)`, ByGroups(NameAttribute, TextWhitespace, Punctuation, TextWhitespace), Push("attr_id")},
				{`\b(n|ne|e|se|s|sw|w|nw|c|_)\b`, NameBuiltin, nil},
				{`\b\D\w*`, NameTag, nil},
				{`[-]?((\.[0-9]+)|([0-9]+(\.[0-9]*)?))`, LiteralNumber, nil},
				{`"(\\"|[^"])*?"`, NameTag, nil},
				{`<`, Punctuation, Push("xml")},
			},
			"attr_id": {
				{`\b\D\w*`, LiteralString, Pop(1)},
				{`[-]?((\.[0-9]+)|([0-9]+(\.[0-9]*)?))`, LiteralNumber, Pop(1)},
				{`"(\\"|[^"])*?"`, LiteralStringDouble, Pop(1)},
				{`<`, Punctuation, Push("#pop", "xml")},
			},
			"xml": {
				{`<`, Punctuation, Push()},
				{`>`, Punctuation, Pop(1)},
				{`\s+`, TextWhitespace, nil},
				{`[^<>\s]`, NameTag, nil},
			},
		}
	},
))
