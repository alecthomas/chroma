package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ErlangShell lexer.
var ErlangShell = internal.Register(MustNewLexer(
	&Config{
		Name:      "Erlang erl session",
		Aliases:   []string{"erl"},
		Filenames: []string{"*.erl-sh"},
		MimeTypes: []string{"text/x-erl-shellsession"},
	},
	Rules{
		"root": {},
	},
))
