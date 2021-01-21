package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RawToken lexer.
var RawToken = internal.Register(MustNewLexer(
	&Config{
		Name:      "Raw token data",
		Aliases:   []string{"raw"},
		MimeTypes: []string{"application/x-pygments-tokens"},
	},
	Rules{
		"root": {},
	},
))
