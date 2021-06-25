package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TrafficScript lexer. For `Riverbed Stingray Traffic Manager
// <http://www.riverbed.com/stingray>`
var TrafficScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "TrafficScript",
		Aliases:   []string{"rts", "trafficscript"},
		Filenames: []string{"*.rts"},
	},
	Rules{
		"root": {},
	},
))
