package chroma

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewlineAtEndOfFile(t *testing.T) {
	l := Coalesce(MustNewLexer(&Config{EnsureNL: true}, Rules{
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err := l.Tokenise(nil, `hello`)
	assert.NoError(t, err)
	assert.Equal(t, []*Token{{Keyword, "hello"}, {Whitespace, "\n"}}, it.Tokens())

	l = Coalesce(MustNewLexer(nil, Rules{
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err = l.Tokenise(nil, `hello`)
	assert.NoError(t, err)
	assert.Equal(t, []*Token{{Error, "hello"}}, it.Tokens())
}
