package chroma

import (
	"iter"
	"slices"
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

type countingLexer struct {
	Lexer
	tokens   []Token
	consumed int
	texts    []string
}

func (l *countingLexer) Tokenise(_ *TokeniseOptions, text string) (iter.Seq[Token], error) {
	l.texts = append(l.texts, text)
	return func(yield func(Token) bool) {
		for _, token := range l.tokens {
			l.consumed++
			if !yield(token) {
				return
			}
		}
	}, nil
}

func makeDelegationTestLexers(t *testing.T) (lang Lexer, root Lexer) {
	t.Helper()
	return mustNewLexer(t, nil, Rules{ // nolint: forbidigo
			"root": {
				{`\<\?`, CommentPreproc, Push("inside")},
				{`.`, Other, nil},
			},
			"inside": {
				{`\?\>`, CommentPreproc, Pop(1)},
				{`\bwhat\b`, Keyword, nil},
				{`\s+`, Whitespace, nil},
			},
		}),
		mustNewLexer(t, nil, Rules{ // nolint: forbidigo
			"root": {
				{`\bhello\b`, Keyword, nil},
				{`\b(world|there)\b`, Name, nil},
				{`\s+`, Whitespace, nil},
			},
		})
}

func TestDelegate(t *testing.T) {
	testdata := []struct {
		name     string
		source   string
		expected []Token
	}{
		{"SourceInMiddle", `hello world <? what ?> there`, []Token{
			{Keyword, "hello"},
			{TextWhitespace, " "},
			{Name, "world"},
			{TextWhitespace, " "},
			// lang
			{CommentPreproc, "<?"},
			{Whitespace, " "},
			{Keyword, "what"},
			{Whitespace, " "},
			{CommentPreproc, "?>"},
			// /lang
			{TextWhitespace, " "},
			{Name, "there"},
		}},
		{"SourceBeginning", `<? what ?> hello world there`, []Token{
			{CommentPreproc, "<?"},
			{TextWhitespace, " "},
			{Keyword, "what"},
			{TextWhitespace, " "},
			{CommentPreproc, "?>"},
			{TextWhitespace, " "},
			{Keyword, "hello"},
			{TextWhitespace, " "},
			{Name, "world"},
			{TextWhitespace, " "},
			{Name, "there"},
		}},
		{"SourceEnd", `hello world <? what there`, []Token{
			{Keyword, "hello"},
			{TextWhitespace, " "},
			{Name, "world"},
			{TextWhitespace, " "},
			// lang
			{CommentPreproc, "<?"},
			{Whitespace, " "},
			{Keyword, "what"},
			{TextWhitespace, " "},
			{Error, "there"},
		}},
		{"SourceMultiple", "hello world <? what ?> hello there <? what ?> hello", []Token{
			{Keyword, "hello"},
			{TextWhitespace, " "},
			{Name, "world"},
			{TextWhitespace, " "},
			{CommentPreproc, "<?"},
			{TextWhitespace, " "},
			{Keyword, "what"},
			{TextWhitespace, " "},
			{CommentPreproc, "?>"},
			{TextWhitespace, " "},
			{Keyword, "hello"},
			{TextWhitespace, " "},
			{Name, "there"},
			{TextWhitespace, " "},
			{CommentPreproc, "<?"},
			{TextWhitespace, " "},
			{Keyword, "what"},
			{TextWhitespace, " "},
			{CommentPreproc, "?>"},
			{TextWhitespace, " "},
			{Keyword, "hello"},
		}},
	}
	lang, root := makeDelegationTestLexers(t)
	delegate := DelegatingLexer(root, lang)
	for _, test := range testdata {
		t.Run(test.name, func(t *testing.T) {
			it, err := delegate.Tokenise(nil, test.source)
			assert.NoError(t, err)
			actual := slices.Collect(it)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestDelegateStreamsRootTokens(t *testing.T) {
	root := &countingLexer{tokens: []Token{
		{Name, "a"},
		{Text, "c"},
		{Name, "e"},
	}}
	language := &countingLexer{tokens: []Token{
		{Other, "a"},
		{Keyword, "b"},
		{Other, "c"},
		{Keyword, "d"},
		{Other, "e"},
	}}
	it, err := DelegatingLexer(root, language).Tokenise(nil, "abcde")
	assert.NoError(t, err)
	assert.Equal(t, 0, root.consumed)
	assert.Equal(t, 5, language.consumed)

	var actual Token
	for token := range it {
		actual = token
		break
	}
	assert.Equal(t, Token{Name, "a"}, actual)
	assert.Equal(t, 2, root.consumed)
	assert.Equal(t, 7, language.consumed)
}

func TestDelegateSplitsRootTokens(t *testing.T) {
	root := &countingLexer{tokens: []Token{{Text, "before  after"}}}
	language := &countingLexer{tokens: []Token{
		{Other, "before "},
		{Keyword, "embedded"},
		{Other, " after"},
	}}
	it, err := DelegatingLexer(root, language).Tokenise(nil, "before embedded after")
	assert.NoError(t, err)
	assert.Equal(t, []string{"before  after"}, root.texts)
	assert.Equal(t, []Token{
		{Text, "before "},
		{Keyword, "embedded"},
		{Text, " after"},
	}, slices.Collect(it))
}
