package chroma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenTypeClassifiers(t *testing.T) {
	require.True(t, GenericDeleted.InCategory(Generic))
	require.True(t, LiteralStringBacktick.InSubCategory(String))
	require.Equal(t, LiteralStringBacktick.String(), "LiteralStringBacktick")
}

func TestSimpleLexer(t *testing.T) {
	lexer, err := NewLexer(
		&Config{
			Name:      "INI",
			Aliases:   []string{"ini", "cfg"},
			Filenames: []string{"*.ini", "*.cfg"},
		},
		map[string][]Rule{
			"root": []Rule{
				{`\s+`, Whitespace, nil},
				{`;.*?$`, Comment, nil},
				{`\[.*?\]$`, Keyword, nil},
				{`(.*?)(\s*)(=)(\s*)(.*?)$`, ByGroups(Name, Whitespace, Operator, Whitespace, String), nil},
			},
		},
	)
	require.NoError(t, err)
	actual, err := lexer.Tokenise(`
	; this is a comment
	[section]
	a = 10
`)
	require.NoError(t, err)
	expected := []Token{
		Token{Whitespace, "\n\t"},
		Token{Comment, "; this is a comment"},
		Token{Whitespace, "\n\t"},
		Token{Keyword, "[section]"},
		Token{Whitespace, "\n\t"},
		Token{Name, "a"},
		Token{Whitespace, " "},
		Token{Operator, "="},
		Token{Whitespace, " "},
		Token{LiteralString, "10"},
		Token{Whitespace, "\n"},
	}
	require.Equal(t, expected, actual)
}
