package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// IrcLogs lexer.
var IrcLogs = internal.Register(MustNewLexer(
	&Config{
		Name:      "IRC Logs",
		Aliases:   []string{"irc"},
		Filenames: []string{"*.weechatlog"},
		MimeTypes: []string{"text/x-irclog"},
	},
	Rules{
		"root": {},
	},
))
