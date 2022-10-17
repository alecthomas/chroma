package chroma

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestCoalesce(t *testing.T) {
	lexer := Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": []Rule{
			{`[!@#$%^&*()]`, Punctuation, nil},
		},
	}))
	actual, err := Tokenise(lexer, nil, "!@#$")
	assert.NoError(t, err)
	expected := []Token{{Punctuation, "!@#$"}}
	assert.Equal(t, expected, actual)
}
