package p

import (
	. "github.com/colin3dmax/chroma" // nolint
	"github.com/colin3dmax/chroma/lexers/internal"
)

var Plaintext = internal.Register(MustNewLexer(
	&Config{
		Name:      "plaintext",
		Aliases:   []string{"text", "plain", "no-highlight"},
		Filenames: []string{"*.txt"},
		MimeTypes: []string{"text/plain"},
	},
	internal.PlaintextRules,
))
