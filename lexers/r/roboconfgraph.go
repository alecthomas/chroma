package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RoboconfGraph lexer for Roboconf <http://roboconf.net/en/roboconf.html> graph files.
var RoboconfGraph = internal.Register(MustNewLexer(
	&Config{
		Name:      "Roboconf Graph",
		Aliases:   []string{"roboconf-graph"},
		Filenames: []string{"*.graph"},
	},
	Rules{
		"root": {},
	},
))
