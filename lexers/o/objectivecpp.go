package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ObjectiveCPP lexer.
var ObjectiveCPP = internal.Register(MustNewLexer(
	&Config{
		Name:      "Objective-C++",
		Aliases:   []string{"objective-c++", "objectivec++", "obj-c++", "objc++"},
		Filenames: []string{"*.mm", "*.hh"},
		MimeTypes: []string{"text/x-objective-c++"},
		// Lower than C++.
		Priority: 0.05,
	},
	Rules{
		"root": {},
	},
))
