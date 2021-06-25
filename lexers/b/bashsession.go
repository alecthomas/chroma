package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// BashSession lexer.
var BashSession = internal.Register(MustNewLexer(
	&Config{
		Name:      "Bash Session",
		Aliases:   []string{"console", "shell-session"},
		Filenames: []string{"*.sh-session", "*.shell-session"},
		MimeTypes: []string{"application/x-shell-session", "application/x-sh-session"},
	},
	Rules{
		"root": {},
	},
))
