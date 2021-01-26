package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Perl6 lexer.
var Perl6 = internal.Register(MustNewLexer(
	&Config{
		Name:    "Perl6",
		Aliases: []string{"perl6", "pl6", "raku"},
		Filenames: []string{"*.pl", "*.pm", "*.nqp", "*.p6", "*.6pl", "*.p6l", "*.pl6",
			"*.6pm", "*.p6m", "*.pm6", "*.t", "*.raku", "*.rakumod", "*.rakutest", "*.rakudoc"},
		MimeTypes: []string{"text/x-perl6", "application/x-perl6"},
	},
	Rules{
		"root": {},
	},
))
