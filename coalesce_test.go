package chroma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCoalesce(t *testing.T) {
	lexer, err := Coalesce(MustNewLexer(nil, Rules{
		"root": []Rule{
			Rule{`[[:punct:]]`, Punctuation, nil},
		},
	}))
	require.NoError(t, err)
	actual, err := lexer.Tokenise("!@#$%")
	require.NoError(t, err)
	expected := []Token{
		Token{Punctuation, "!@#$%"},
	}
	require.Equal(t, expected, actual)
}
