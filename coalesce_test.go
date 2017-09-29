package chroma

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestCoalesce(t *testing.T) {
	lexer := Coalesce(MustNewLexer(&Config{DontEnsureNL: true}, Rules{
		"root": []Rule{
			{`[!@#$%^&*()]`, Punctuation, nil},
		},
	}))
	actual, err := Tokenise(lexer, nil, "!@#$")
	assert.NoError(t, err)
	expected := []*Token{{Punctuation, "!@#$"}}
	assert.Equal(t, expected, actual)
}
