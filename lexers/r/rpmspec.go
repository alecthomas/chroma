package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RPMSpec lexer.
var RPMSpec = internal.Register(MustNewLexer(
	&Config{
		Name:      "RPMSpec",
		Aliases:   []string{"spec"},
		Filenames: []string{"*.spec"},
		MimeTypes: []string{"text/x-rpm-spec"},
	},
	Rules{
		"root": {},
	},
))
