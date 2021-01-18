package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Todotxt lexer. Lexer for Todo.txt <http://todotxt.com/> todo list format.
var Todotxt = internal.Register(MustNewLexer(
	&Config{
		Name:    "Todotxt",
		Aliases: []string{"todotxt"},
		// *.todotxt is not a standard extension for Todo.txt files; including it
		// makes testing easier, and also makes autodetecting file type easier.
		Filenames: []string{"todo.txt", "*.todotxt"},
		MimeTypes: []string{"text/x-todo"},
	},
	Rules{
		"root": {},
	},
))
