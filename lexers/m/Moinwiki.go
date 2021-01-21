package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MoinWiki lexer.
var MoinWiki = internal.Register(MustNewLexer(
	&Config{
		Name:      "MoinMoin/Trac Wiki markup",
		Aliases:   []string{"trac-wiki", "moin"},
		MimeTypes: []string{"text/x-trac-wiki"},
	},
	Rules{
		"root": {},
	},
))
