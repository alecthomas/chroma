package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ElixirConsole lexer.
var ElixirConsole = internal.Register(MustNewLexer(
	&Config{
		Name:      "Elixir iex session",
		Aliases:   []string{"iex"},
		MimeTypes: []string{"text/x-elixir-shellsession"},
	},
	Rules{
		"root": {},
	},
))
