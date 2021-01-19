package k

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Kconfig lexer.
var Kconfig = internal.Register(MustNewLexer(
	&Config{
		Name:    "Kconfig",
		Aliases: []string{"kconfig", "menuconfig", "linux-config", "kernel-config"},
		// Adjust this if new kconfig file names appear in your environment.
		Filenames: []string{"Kconfig*", "*Config.in*", "external.in*", "standard-modules.in"},
		MimeTypes: []string{"text/x-kconfig"},
	},
	Rules{
		"root": {},
	},
))
