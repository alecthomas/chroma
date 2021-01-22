package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RQL lexer for Relation Query Language <http://www.logilab.org/project/rql>
var RQL = internal.Register(MustNewLexer(
	&Config{
		Name:      "RQL",
		Aliases:   []string{"rql"},
		Filenames: []string{"*.rql"},
		MimeTypes: []string{"text/x-rql"},
	},
	Rules{
		"root": {},
	},
))
