package lexers

import (
	"strings"

	. "github.com/alecthomas/chroma/v3" // nolint
)

// Blade lexer is Laravel's Blade templating language embedded in HTML.
var Blade = Register(DelegatingLexer(HTML, MustNewXMLLexer(
	embedded,
	"embedded/blade.xml",
).SetConfig(
	&Config{
		Name:      "Blade",
		Aliases:   []string{"blade", "html+blade"},
		Filenames: []string{"*.blade.php"},
		MimeTypes: []string{"text/x-blade"},
		DotAll:    true,
		Priority:  3,
	},
).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, "@endforeach") || strings.Contains(text, "@endif") || strings.Contains(text, "@extends(") {
		return 0.4
	}
	return 0.0
})))
