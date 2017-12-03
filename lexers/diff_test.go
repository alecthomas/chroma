package lexers

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
)

func TestDiffLexerWithoutTralingNewLine(t *testing.T) {
	diffLexer := Get("diff")
	it, err := diffLexer.Tokenise(nil, "-foo\n+bar")
	assert.NoError(t, err)
	actual := it.Tokens()
	expected := []*chroma.Token{
		&chroma.Token{chroma.GenericDeleted, "-foo\n"},
		&chroma.Token{chroma.GenericInserted, "+bar\n"},
	}
	assert.Equal(t, expected, actual)
}
