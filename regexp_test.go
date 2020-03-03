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
	assert.Equal(t, []Token{{Keyword, "hello"}, {Whitespace, "\n"}}, it.Tokens())

	l = Coalesce(MustNewLexer(nil, Rules{
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err = l.Tokenise(nil, `hello`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Error, "hello"}}, it.Tokens())
}

func TestMatchingAtStart(t *testing.T) {
	l := Coalesce(MustNewLexer(&Config{}, Rules{
		"root": {
			{`\s+`, Whitespace, nil},
			{`^-`, Punctuation, Push("directive")},
			{`->`, Operator, nil},
		},
		"directive": {
			{"module", NameEntity, Pop(1)},
		},
	}))
	it, err := l.Tokenise(nil, `-module ->`)
	assert.NoError(t, err)
	assert.Equal(t,
		[]Token{{Punctuation, "-"}, {NameEntity, "module"}, {Whitespace, " "}, {Operator, "->"}},
		it.Tokens())
}

func TestEnsureLFOption(t *testing.T) {
	l := Coalesce(MustNewLexer(&Config{}, Rules{
		"root": {
			{`(\w+)(\r?\n|\r)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err := l.Tokenise(&TokeniseOptions{
		State:    "root",
		EnsureLF: true,
	}, "hello\r\nworld\r")
	assert.NoError(t, err)
	assert.Equal(t, []Token{
		{Keyword, "hello"},
		{Whitespace, "\n"},
		{Keyword, "world"},
		{Whitespace, "\n"},
	}, it.Tokens())

	l = Coalesce(MustNewLexer(nil, Rules{
		"root": {
			{`(\w+)(\r?\n|\r)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err = l.Tokenise(&TokeniseOptions{
		State:    "root",
		EnsureLF: false,
	}, "hello\r\nworld\r")
	assert.NoError(t, err)
	assert.Equal(t, []Token{
		{Keyword, "hello"},
		{Whitespace, "\r\n"},
		{Keyword, "world"},
		{Whitespace, "\r"},
	}, it.Tokens())
}

func TestEnsureLFFunc(t *testing.T) {
	tests := []struct{ in, out string }{
		{in: "", out: ""},
		{in: "abc", out: "abc"},
		{in: "\r", out: "\n"},
		{in: "a\r", out: "a\n"},
		{in: "\rb", out: "\nb"},
		{in: "a\rb", out: "a\nb"},
		{in: "\r\n", out: "\n"},
		{in: "a\r\n", out: "a\n"},
		{in: "\r\nb", out: "\nb"},
		{in: "a\r\nb", out: "a\nb"},
		{in: "\r\r\r\n\r", out: "\n\n\n\n"},
	}
	for _, test := range tests {
		out := ensureLF(test.in)
		assert.Equal(t, out, test.out)
	}
}
