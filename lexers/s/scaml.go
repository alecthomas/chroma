package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Scaml lexer. For Scaml markup <http://scalate.fusesource.org/>. Scaml is Haml for Scala.
var Scaml = internal.Register(MustNewLexer(
	&Config{
		Name:      "Scaml",
		Aliases:   []string{"scaml"},
		Filenames: []string{"*.scaml"},
		MimeTypes: []string{"text/x-scaml"},
	},
	Rules{
		"root": {},
	},
))
