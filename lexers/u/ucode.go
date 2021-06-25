package u

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ucode lexer.
var Ucode = internal.Register(MustNewLexer(
	&Config{
		Name:      "ucode",
		Aliases:   []string{"ucode"},
		Filenames: []string{"*.u", "*.u1", "*.u2"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// endsuspend and endrepeat are unique to this language, and
	// \self, /self doesn't seem to get used anywhere else either.
	var result float32

	if strings.Contains(text, "endsuspend") {
		result += 0.1
	}

	if strings.Contains(text, "endrepeat") {
		result += 0.1
	}

	if strings.Contains(text, ":=") {
		result += 0.01
	}

	if strings.Contains(text, "procedure") && strings.Contains(text, "end") {
		result += 0.01
	}

	// This seems quite unique to unicon -- doesn't appear in any other
	// example source we have (A quick search reveals that \SELF appears in
	// Perl/Raku code)
	if strings.Contains(text, `\self`) && strings.Contains(text, "/self") {
		result += 0.5
	}

	return result
}))
