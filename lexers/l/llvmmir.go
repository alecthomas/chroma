package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LlvmMir lexer.
var LlvmMir = internal.Register(MustNewLexer(
	&Config{
		Name:      "LLVM-MIR",
		Aliases:   []string{"llvm-mir"},
		Filenames: []string{"*.mir"},
	},
	Rules{
		"root": {},
	},
))
