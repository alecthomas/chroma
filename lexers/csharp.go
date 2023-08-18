package lexers

import (
	. "github.com/alecthomas/chroma/v2" // nolint
)

// CSharp lexer.
var CSharp = Register(MustNewXMLLexer(
	embedded,
	"embedded/csharp.xml",
).SetConfig(
	&Config{
		Name:      "C#",
		Aliases:   []string{"csharp", "c#", "cs"},
		Filenames: []string{"*.cs"},
		MimeTypes: []string{"text/x-csharp"},
		DotAll:    true,
		EnsureNL:  true,
	},
))
