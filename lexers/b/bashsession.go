package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// BashSession lexer.
var BashSession = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Bash Session",
		Aliases:   []string{"bash-session", "console", "shell-session"},
		Filenames: []string{".sh-session"},
		MimeTypes: []string{"application/x-shell-session", "application/x-sh-session", "text/x-sh"},
		EnsureNL:  true,
	},
	bashsessionRules,
))

func bashsessionRules() Rules {
	return Rules{
		"root": {
			{`(^[#$%>]\s*)(.*\n?)`, ByGroups(GenericPrompt, Using(Bash)), nil},
			{`^.+\n?`, GenericOutput, nil},
		},
	}
}
