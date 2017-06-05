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
			"root": {
				{`\s+`, Whitespace, nil},
				{`;.*?$`, Comment, nil},
				{`\[.*?\]$`, Keyword, nil},
				{`(.*?)(\s*)(=)(\s*)(.*?)$`, ByGroups(Name, Whitespace, Operator, Whitespace, String), nil},
			},
		},
	)
	require.NoError(t, err)
	actual, err := Tokenise(lexer, nil, `
	; this is a comment
	[section]
	a = 10
`)
	require.NoError(t, err)
	expected := []*Token{
		{Whitespace, "\n\t"},
		{Comment, "; this is a comment"},
		{Whitespace, "\n\t"},
		{Keyword, "[section]"},
		{Whitespace, "\n\t"},
		{Name, "a"},
		{Whitespace, " "},
		{Operator, "="},
		{Whitespace, " "},
		{LiteralString, "10"},
		{Whitespace, "\n"},
	}
	require.Equal(t, expected, actual)
}
