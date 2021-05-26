package chroma

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestCoalesce(t *testing.T) {
	lexer := Coalesce(MustNewLexer(nil, Rules{ // nolint: forbidigo
		"root": []Rule{
			{`[!@#$%^&*()]`, Punctuation, nil},
		},
	}), false)
	actual, err := Tokenise(lexer, nil, "!@#$")
	assert.NoError(t, err)
	expected := []Token{{Punctuation, "!@#$"}}
	assert.Equal(t, expected, actual)
}

func TestCoalesceKeepLineSplits(t *testing.T) {
	lexer := Coalesce(MustNewLexer(nil, Rules{ // nolint: forbidigo
		"root": []Rule{
			{`.*\n`, Text, nil},
		},
	}), true)
	actual, err := Tokenise(lexer, nil, "foo\nbar\n")
	assert.NoError(t, err)
	expected := []Token{{Text, "foo\n"}, {Text, "bar\n"}}
	assert.Equal(t, expected, actual)
}
