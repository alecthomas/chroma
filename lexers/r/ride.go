package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ride lexer. For Ride <https://docs.wavesplatform.com/en/ride/about-ride.html>
// source code.
var Ride = internal.Register(MustNewLexer(
	&Config{
		Name:      "Ride",
		Aliases:   []string{"ride"},
		Filenames: []string{"*.ride"},
		MimeTypes: []string{"text/x-ride"},
	},
	Rules{
		"root": {},
	},
))
