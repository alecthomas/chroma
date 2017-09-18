package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Diff lexer.
var Diff = Register(MustNewLexer(
	&Config{
		Name:      "Diff",
		Aliases:   []string{"diff", "udiff"},
		Filenames: []string{"*.diff", "*.patch"},
		MimeTypes: []string{"text/x-diff", "text/x-patch"},
	},
	Rules{
		"root": {
			{` .*\n`, Text, nil},
			{`\+.*\n`, GenericInserted, nil},
			{`-.*\n`, GenericDeleted, nil},
			{`!.*\n`, GenericStrong, nil},
			{`@.*\n`, GenericSubheading, nil},
			{`([Ii]ndex|diff).*\n`, GenericHeading, nil},
			{`=.*\n`, GenericHeading, nil},
			{`.*\n`, Text, nil},
		},
	},
))
