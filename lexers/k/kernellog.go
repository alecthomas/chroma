package k

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// KernelLog lexer.
var KernelLog = internal.Register(MustNewLexer(
	&Config{
		Name:      "Kernel log",
		Aliases:   []string{"kmsg", "dmesg"},
		Filenames: []string{"*.kmsg", "*.dmesg"},
	},
	Rules{
		"root": {},
	},
))
