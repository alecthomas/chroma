package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RubyIRBSession lexer. For Ruby interactive console (**irb**) output.
var RubyIRBSession = internal.Register(MustNewLexer(
	&Config{
		Name:      "Ruby irb session",
		Aliases:   []string{"rbcon", "irb"},
		MimeTypes: []string{"text/x-ruby-shellsession"},
	},
	Rules{
		"root": {},
	},
))
