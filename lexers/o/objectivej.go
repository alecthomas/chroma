package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Objective-J lexer.
var ObjectiveJ = internal.Register(MustNewLexer(
	&Config{
		Name:      "Objective-J",
		Aliases:   []string{"objective-j", "objectivej", "obj-j", "objj"},
		Filenames: []string{"*.j"},
		MimeTypes: []string{"text/x-objective-j"},
	},
	Rules{
		"root": {},
	},
))
