package chroma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoalesce(t *testing.T) {
	lexer := Coalesce(MustNewLexer(nil, Rules{ // nolint: forbidigo
		"root": []Rule{
			{`[!@#$%^&*()]`, Punctuation, nil},
		},
	}))
	actual, err := Tokenise(lexer, nil, "!@#$")
	assert.NoError(t, err)
	expected := []Token{{Punctuation, "!@#$"}}
	assert.Equal(t, expected, actual)
}
