package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Genshi Text lexer.
var GenshiTemplate = Register(MustNewLexer(
	&Config{
		Name:      "Genshi Text",
		Aliases:   []string{"genshitext"},
		Filenames: []string{},
		MimeTypes: []string{"application/x-genshi-text", "text/x-genshi"},
	},
	Rules{
		"root": {
			{`[^#$\s]+`, Other, nil},
			{`^(\s*)(##.*)$`, ByGroups(Text, Comment), nil},
			{`^(\s*)(#)`, ByGroups(Text, CommentPreproc), Push("directive")},
			Include("variable"),
			{`[#$\s]`, Other, nil},
		},
		"directive": {
			{`\n`, Text, Pop(1)},
			{`(?:def|for|if)\s+.*`, Using(Python, nil), Pop(1)},
			{`(choose|when|with)([^\S\n]+)(.*)`, ByGroups(Keyword, Text, Using(Python, nil)), Pop(1)},
			{`(choose|otherwise)\b`, Keyword, Pop(1)},
			{`(end\w*)([^\S\n]*)(.*)`, ByGroups(Keyword, Text, Comment), Pop(1)},
		},
		"variable": {
			{`(?<!\$)(\$\{)(.+?)(\})`, ByGroups(CommentPreproc, Using(Python, nil), CommentPreproc), nil},
			{`(?<!\$)(\$)([a-zA-Z_][\w.]*)`, NameVariable, nil},
		},
	},
))
