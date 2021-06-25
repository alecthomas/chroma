package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RNCCompact lexer. For RelaxNG-compact <http://relaxng.org> syntax.
var RNCCompact = internal.Register(MustNewLexer(
	&Config{
		Name:      "Relax-NG Compact",
		Aliases:   []string{"rnc", "rng-compact"},
		Filenames: []string{"*.rnc"},
	},
	Rules{
		"root": {},
	},
))
