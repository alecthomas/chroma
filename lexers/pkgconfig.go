package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Pkgconfig lexer.
var Pkgconfig = Register(MustNewLexer(
	&Config{
		Name:      "PkgConfig",
		Aliases:   []string{"pkgconfig"},
		Filenames: []string{"*.pc"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {
			{`#.*$`, CommentSingle, nil},
			{`^(\w+)(=)`, ByGroups(NameAttribute, Operator), nil},
			{`^([\w.]+)(:)`, ByGroups(NameTag, Punctuation), Push("spvalue")},
			Include("interp"),
			{`[^${}#=:\n.]+`, Text, nil},
			{`.`, Text, nil},
		},
		"interp": {
			{`\$\$`, Text, nil},
			{`\$\{`, LiteralStringInterpol, Push("curly")},
		},
		"curly": {
			{`\}`, LiteralStringInterpol, Pop(1)},
			{`\w+`, NameAttribute, nil},
		},
		"spvalue": {
			Include("interp"),
			{`#.*$`, CommentSingle, Pop(1)},
			{`\n`, Text, Pop(1)},
			{`[^${}#\n]+`, Text, nil},
			{`.`, Text, nil},
		},
	},
))
