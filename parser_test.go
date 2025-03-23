package chroma

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestParser(t *testing.T) {
	ast, err := parser.ParseString("", `
config {
	name "INI"
	aliases "ini", "cfg"
	filenames "*.ini", "*.cfg", "*.inf", "*.service", "*.socket", ".gitconfig",
			  ".editorconfig", "pylintrc", ".pylintrc"
	mime-types "text/x-ini", "text/inf"
	priority 0.1
}

state root {
	/\s+/ text
	/[;#].*/ commentsingle
	/\[.*?\]$/ keyword
	/(.*?)([ \t]*)(=)([ \t]*)(.*(?:\n[ \t].+)*)/ by groups
		nameattribute, text, operator, text, literalstring
	/(.+?)$/ nameattribute
}
`)
	assert.NoError(t, err)
	assert.Equal(t, &AST{
		Config: &Config{
			Name: "INI",
			Aliases: []string{
				"ini",
				"cfg",
			},
			Filenames: []string{
				"*.ini",
				"*.cfg",
				"*.inf",
				"*.service",
				"*.socket",
				".gitconfig",
				".editorconfig",
				"pylintrc",
				".pylintrc",
			},
			MimeTypes: []string{
				"text/x-ini",
				"text/inf",
			},
			Priority: 0.1,
		},
		States: []stateAST{
			{Name: "root",
				Rules: []Rule{
					{Pattern: `\s+`, Type: &tokenTypeAST{Text}},
					{Pattern: `[;#].*`, Type: &tokenTypeAST{CommentSingle}},
					{Pattern: `\[.*?\]$`, Type: &tokenTypeAST{Keyword}},
					{Pattern: `(.*?)([ \t]*)(=)([ \t]*)(.*(?:\n[ \t].+)*)`, Type: &byGroupsEmitter{Emitters{
						&tokenTypeAST{NameAttribute},
						&tokenTypeAST{Text},
						&tokenTypeAST{Operator},
						&tokenTypeAST{Text},
						&tokenTypeAST{LiteralString},
					}}},
					{Pattern: `(.+?)$`, Type: &tokenTypeAST{NameAttribute}},
				},
			},
		},
	}, ast)
}
