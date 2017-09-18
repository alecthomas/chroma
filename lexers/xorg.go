package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Xorg lexer.
var Xorg = Register(MustNewLexer(
	&Config{
		Name:      "Xorg",
		Aliases:   []string{"xorg.conf"},
		Filenames: []string{"xorg.conf"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {
			{`\s+`, Text, nil},
			{`#.*$`, Comment, nil},
			{`((|Sub)Section)(\s+)("\w+")`, ByGroups(LiteralStringEscape, LiteralStringEscape, Text, LiteralStringEscape), nil},
			{`(End(|Sub)Section)`, LiteralStringEscape, nil},
			{`(\w+)(\s+)([^\n#]+)`, ByGroups(NameBuiltin, Text, NameConstant), nil},
		},
	},
))
