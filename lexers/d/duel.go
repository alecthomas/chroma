package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Duel lexer.
var Duel = internal.Register(MustNewLexer(
	&Config{
		Name:      "Duel",
		Aliases:   []string{"duel", "jbst", "jsonml+bst"},
		Filenames: []string{"*.duel", "*.jbst"},
		MimeTypes: []string{"text/x-duel", "text/x-jbst"},
	},
	Rules{
		"root": {},
	},
))
