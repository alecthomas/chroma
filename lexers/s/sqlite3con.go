package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Sqlite3con lexer. Lexer for example sessions using sqlite3.
var Sqlite3con = internal.Register(MustNewLexer(
	&Config{
		Name:      "sqlite3con",
		Aliases:   []string{"sqlite3"},
		Filenames: []string{"*.sqlite3-console"},
		MimeTypes: []string{"text/x-sqlite3-console"},
	},
	Rules{
		"root": {},
	},
))
