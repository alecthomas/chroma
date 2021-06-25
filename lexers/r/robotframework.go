package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// RobotFramework lexer for Robot Framework <http://robotframework.org> test data.
var RobotFramework = internal.Register(MustNewLexer(
	&Config{
		Name:      "RobotFramework",
		Aliases:   []string{"robotframework"},
		Filenames: []string{"*.robot"},
		MimeTypes: []string{"text/x-robotframework"},
	},
	Rules{
		"root": {},
	},
))
