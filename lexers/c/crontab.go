package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Crontab lexer.
var Crontab = internal.Register(MustNewLexer(
	&Config{
		Name:      "Crontab",
		Aliases:   []string{"crontab"},
		Filenames: []string{"crontab"},
	},
	Rules{
		"root": {},
	},
))
