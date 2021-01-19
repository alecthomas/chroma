package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// LlvmMirBody lexer.
var LlvmMirBody = internal.Register(MustNewLexer(
	&Config{
		Name:    "LLVM-MIR Body",
		Aliases: []string{"llvm-mir-body"},
	},
	Rules{
		"root": {},
	},
))
