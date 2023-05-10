package chroma

import (
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

type mockReader struct {
	reader    strings.Reader
	t         *testing.T
	blockSize int
}

func (r *mockReader) Read(b []byte) (int, error) {
	r.t.Helper()
	assert.True(r.t, len(b) <= r.blockSize)
	return r.reader.Read(b)
}

func testAssertReader(t *testing.T, lexer Lexer, options *TokeniseOptions, text string, blockSize int) (Iterator, error) {
	t.Helper()

	reader := &mockReader{
		reader: *strings.NewReader(text),
		t:      t,
	}
	reader.blockSize = blockSize

	return lexer.TokeniseStream(options, reader, blockSize, int(reader.reader.Size()))
}

func mustNewLexer(t *testing.T, config *Config, rules Rules) *RegexLexer { // nolint: forbidigo
	lexer, err := NewLexer(config, func() Rules {
		return rules
	})
	assert.NoError(t, err)
	return lexer
}

func TestNewlineAtEndOfFile(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{EnsureNL: true}, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err := l.Tokenise(nil, `hello`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Keyword, "hello"}, {Whitespace, "\n"}}, it.Tokens())

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err = l.Tokenise(nil, `hello`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Error, "hello"}}, it.Tokens())
}

func TestNewlineAtEndOfFileStream(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{EnsureNL: true}, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))

	it, err := testAssertReader(t, l, nil, `hello`, 128)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Keyword, "hello"}, {Whitespace, "\n"}}, it.Tokens())

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err = testAssertReader(t, l, nil, `hello`, 128)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Error, "hello"}}, it.Tokens())
}

func TestMatchingAtStart(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
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

func TestMatchingAtStartStream(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`\s+`, Whitespace, nil},
			{`^-`, Punctuation, Push("directive")},
			{`->`, Operator, nil},
		},
		"directive": {
			{"module", NameEntity, Pop(1)},
		},
	}))

	it, err := testAssertReader(t, l, nil, `-module ->`, 6)
	assert.NoError(t, err)
	assert.Equal(t,
		[]Token{{Punctuation, "-"}, {NameEntity, "module"}, {Whitespace, " "}, {Operator, "->"}},
		it.Tokens())

	it, err = testAssertReader(t, l, nil, `-module ->`, 7)
	assert.NoError(t, err)
	assert.Equal(t,
		[]Token{{Punctuation, "-"}, {NameEntity, "module"}, {Whitespace, " "}, {Operator, "->"}},
		it.Tokens())

	it, err = testAssertReader(t, l, nil, `-module ->`, 3)
	assert.NoError(t, err)
	assert.Equal(t,
		[]Token{{Punctuation, "-"}, {Error, "module ->"}},
		it.Tokens())
}

func TestEnsureLFOption(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
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

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
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

func TestByGroupNames(t *testing.T) {
	l := Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)(?<operator>=)(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:      String,
					`operator`: Operator,
					`value`:    String,
				}),
				nil,
			},
		},
	}))
	it, err := l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc`}, {Operator, `=`}, {String, `123`}}, it.Tokens())

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)(?<operator>=)(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:   String,
					`value`: String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc`}, {Error, `=`}, {String, `123`}}, it.Tokens())

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)=(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:   String,
					`value`: String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc123`}}, it.Tokens())

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)(?<op>=)(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:      String,
					`operator`: Operator,
					`value`:    String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc`}, {Error, `=`}, {String, `123`}}, it.Tokens())

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`\w+=\w+`,
				ByGroupNames(map[string]Emitter{
					`key`:      String,
					`operator`: Operator,
					`value`:    String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Error, `abc=123`}}, it.Tokens())
}
