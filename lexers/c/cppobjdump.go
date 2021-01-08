package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CppObjdump lexer.
var CppObjdump = internal.Register(MustNewLexer(
	&Config{
		Name:      "cpp-objdump",
		Aliases:   []string{"cpp-objdump", "c++-objdumb", "cxx-objdump"},
		Filenames: []string{"*.cpp-objdump", "*.c++-objdump", "*.cxx-objdump"},
		MimeTypes: []string{"text/x-cpp-objdump"},
	},
	Rules{
		"root": {},
	},
))
