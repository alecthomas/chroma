package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/h"
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/lexers/t"
)

// Svelte lexer.
var Svelte = internal.Register(DelegatingLexer(h.HTML, MustNewLazyLexer(
	&Config{
		Name:      "Svelte",
		Aliases:   []string{"svelte"},
		Filenames: []string{"*.svelte"},
		MimeTypes: []string{"application/x-svelte"},
		DotAll:    true,
	},
	svelteRules,
)))

func svelteRules() Rules {
	return Rules{
		"root": {
			{`(<\s*script\s*lang\s*=\s*['"](?:ts|typescript)['"]\s*>)(.+?)(<\s*/\s*script\s*>)`, ByGroups(Other, Using(t.TypeScript), Other), nil},
			{`\{`, Punctuation, Push("templates")},
			{`[^{]`, Other, nil},
		},
		"templates": {
			{`}`, Punctuation, Pop(1)},
			{`@(debug|html)\b`, Keyword, nil},
			{`(#|/)(await|each|if)\b`, Keyword, nil},
			{`(:else)(\s+)(if)?\b`, ByGroups(Keyword, Text, Keyword), nil},
			{`:(catch|then)\b`, Keyword, nil},
			{`then\b`, Keyword, nil},
			{`[^}]+`, Using(t.TypeScript), nil},
		},
	}
}
