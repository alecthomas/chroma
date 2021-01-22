package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// NasmObjdump lexer.
var NasmObjdump = internal.Register(MustNewLexer(
	&Config{
		Name:      "objdump-nasm",
		Aliases:   []string{"objdump-nasm"},
		Filenames: []string{"*.objdump-intel"},
		MimeTypes: []string{"text/x-nasm-objdump"},
	},
	Rules{
		"root": {},
	},
))
