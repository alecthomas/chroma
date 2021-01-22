package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RoboconfInstances lexer for Roboconf <http://roboconf.net/en/roboconf.html> instances files.
var RoboconfInstances = internal.Register(MustNewLexer(
	&Config{
		Name:      "Roboconf Instances",
		Aliases:   []string{"roboconf-instances"},
		Filenames: []string{"*.instances"},
	},
	Rules{
		"root": {},
	},
))
