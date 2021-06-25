package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Tiddler lexer. For TiddlyWiki5 <https://tiddlywiki.com/#TiddlerFiles> markup.
var Tiddler = internal.Register(MustNewLexer(
	&Config{
		Name:      "tiddler",
		Aliases:   []string{"tid"},
		Filenames: []string{"*.tid"},
		MimeTypes: []string{"text/vnd.tiddlywiki"},
	},
	Rules{
		"root": {},
	},
))
