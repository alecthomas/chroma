package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Console lexer.
var Console = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Console",
		Aliases:   []string{"console", "shell-session"},
		Filenames: []string{},
		MimeTypes: []string{},
		EnsureNL: true,
	},
	consoleRules,
))

func consoleRules() Rules {
	return Rules{
		"root": {
			{`^((?:\([^)]*\))?(?:\[?[a-zA-Z0-9_.-]+(?:@[a-zA-Z0-9\-]+)?\]?)?(?:(?:[: \t])[^\n]+)?[ \t]*[#$%])([ \t]*.*\n)`, EmitterFunc(consolePrompt), nil},
			{`.+`, GenericOutput, nil},
			{`\n`, GenericOutput, nil},
		},
	}
}

func consolePrompt(groups []string, lexer Lexer) Iterator {
	tokens := []Token{
		{GenericPrompt, groups[1]},
		{Text, groups[2]},
	}
	return Literator(tokens...)
}
