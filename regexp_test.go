package chroma

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustNewLexer(t *testing.T, config *Config, rules Rules) *RegexLexer { // nolint: forbidigo
	lexer, err := NewLexer(config, func() Rules {
		return rules
	})
	require.NoError(t, err)
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
		out, _ := ensureLF(test.in)
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

func TestTokenizeWithOffsets(t *testing.T) {
	type testCase struct {
		name                           string
		input                          string
		expected                       []Token
		expectedOriginalLengths        []int
		expectedOriginalLengthsInRunes []int
	}

	tests := []testCase{
		{
			name:  "empty",
			input: "",
		},
		{
			name:                           "-m ->",
			input:                          "-m ->",
			expected:                       []Token{{Punctuation, "-"}, {NameEntity, "m"}, {Whitespace, " "}, {Operator, "->"}},
			expectedOriginalLengths:        []int{1, 1, 1, 2},
			expectedOriginalLengthsInRunes: []int{1, 1, 1, 2},
		},
		{
			name:                           `-m\r\n-m`,
			input:                          "-m\r\n-m",
			expected:                       []Token{{Punctuation, "-"}, {NameEntity, "m"}, {Whitespace, "\n"}, {Punctuation, "-"}, {NameEntity, "m"}},
			expectedOriginalLengths:        []int{1, 1, 2, 1, 1},
			expectedOriginalLengthsInRunes: []int{1, 1, 2, 1, 1},
		},
		{
			name:                           `-m\r\n\r\n\r\n`,
			input:                          "-m\r\n\r\n\r\n",
			expected:                       []Token{{Punctuation, "-"}, {NameEntity, "m"}, {Whitespace, "\n\n\n"}},
			expectedOriginalLengths:        []int{1, 1, 6},
			expectedOriginalLengthsInRunes: []int{1, 1, 6},
		},
		{
			name:                           `\r\n-m\r\n`,
			input:                          "\r\n-m\r\n",
			expected:                       []Token{{Whitespace, "\n"}, {Punctuation, "-"}, {NameEntity, "m"}, {Whitespace, "\n"}},
			expectedOriginalLengths:        []int{2, 1, 1, 2},
			expectedOriginalLengthsInRunes: []int{2, 1, 1, 2},
		},
		{
			name:                           `\n-m\r\n-m\n`,
			input:                          "\n-m\r\n-m\n",
			expected:                       []Token{{Whitespace, "\n"}, {Punctuation, "-"}, {NameEntity, "m"}, {Whitespace, "\n"}, {Punctuation, "-"}, {NameEntity, "m"}, {Whitespace, "\n"}},
			expectedOriginalLengths:        []int{1, 1, 1, 2, 1, 1, 1},
			expectedOriginalLengthsInRunes: []int{1, 1, 1, 2, 1, 1, 1},
		},
		{
			name:                           `\n  \r\n  ->`,
			input:                          "\n  \r\n  ->",
			expected:                       []Token{{Whitespace, "\n  \n  "}, {Operator, "->"}},
			expectedOriginalLengths:        []int{7, 2},
			expectedOriginalLengthsInRunes: []int{7, 2},
		},
		{
			// Note: the first space in this input is an enspace (U+2002) taking 3 bytes in unicode
			name:                           `\n \r\n  ->`,
			input:                          "\n \r\n  ->",
			expected:                       []Token{{Whitespace, "\n \n  "}, {Operator, "->"}},
			expectedOriginalLengths:        []int{8, 2},
			expectedOriginalLengthsInRunes: []int{6, 2},
		},
	}

	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`\s+`, Whitespace, nil},
			{`^-`, Punctuation, Push("directive")},
			{`->`, Operator, nil},
		},
		"directive": {
			{"m", NameEntity, Pop(1)},
		},
	}))

	lex, ok := l.(TokeniserWithOriginalLen)
	if !ok {
		assert.True(t, ok, "lexer is not a TokenizerWithOffsets")
	}

	type checkCountFunc func(tok *Token, lengthsIter *OriginalLenIterator, tc testCase, tokenIndex int)

	checkLenInBytes := func(tok *Token, lengthsIter *OriginalLenIterator, tc testCase, tokenIndex int) {
		ln := lengthsIter.OriginalLen(tok)
		assert.Equal(t, tc.expectedOriginalLengths[tokenIndex], ln, "testcase %s byte lengths token %d", tc.name, tokenIndex)
	}

	checkLenInRunes := func(tok *Token, lengthsIter *OriginalLenIterator, tc testCase, tokenIndex int) {
		ln, err := lengthsIter.OriginalLenRunes(tok)
		assert.NoError(t, err)
		assert.Equal(t, tc.expectedOriginalLengthsInRunes[tokenIndex], ln, "testcase %s rune lengths token %d", tc.name, tokenIndex)
	}

	checkCountTypes := []checkCountFunc{
		checkLenInBytes,
		checkLenInRunes,
	}

	for _, checkCountType := range checkCountTypes {
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				tc := tc

				it, lengthsIter, err := lex.TokeniseWithOriginalLen(nil, tc.input)
				assert.NoError(t, err)

				for i := 0; ; i++ {
					tok := it()
					if tok == EOF {
						break
					}

					assert.Less(t, i, len(tc.expected))
					assert.Equal(t, tc.expected[i], tok)

					checkCountType(&tok, &lengthsIter, tc, i)
				}
			})
		}
	}
}

func TestTokenizeWithOffsetsForDelegatingLexer(t *testing.T) {
	outerLexer := mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`\s+`, Whitespace, nil},
			{`\d+`, Other, nil},
			{`->`, Operator, nil},
		},
	})

	innerLexer := mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`[0-5]+`, NumberInteger, nil},
			{`[6-9]+`, NumberHex, nil},
		},
	})

	input := "->\r\n2399"
	expected := []Token{{Operator, "->"}, {Whitespace, "\n"}, {LiteralNumberInteger, "23"}, {LiteralNumberHex, "99"}}
	expectedOriginalLengths := []int{2, 2, 2, 2}

	checkLenInBytes := func(tok *Token, lengthsIter *OriginalLenIterator, tokenIndex int) {
		ln := lengthsIter.OriginalLen(tok)
		assert.Equal(t, expectedOriginalLengths[tokenIndex], ln, "byte lengths token %d", tokenIndex)
	}

	l := DelegatingLexer(innerLexer, outerLexer)

	lex, ok := l.(TokeniserWithOriginalLen)
	if !ok {
		assert.True(t, ok, "lexer is not a TokenizerWithOffsets")
	}

	it, lengthsIter, err := lex.TokeniseWithOriginalLen(nil, input)
	assert.NoError(t, err)

	for i := 0; ; i++ {

		tok := it()
		if tok == EOF {
			break
		}

		assert.Less(t, i, len(expected))
		assert.Equal(t, expected[i], tok)

		checkLenInBytes(&tok, &lengthsIter, i)
	}
}
