package chroma

import (
	"testing"

	"github.com/alecthomas/assert"
)

var (
	delegateSourceMiddle = `hello world <? what ?> there`
	delegateSourceEnd    = `hello world <? what there`
)

func makeDelegationTestLexers() (lang Lexer, root Lexer) {
	return MustNewLexer(nil, Rules{
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
		MustNewLexer(nil, Rules{
			"root": {
				{`\bhello\b`, Keyword, nil},
				{`\b(world|there)\b`, Name, nil},
				{`\s+`, Whitespace, nil},
			},
		})
}

func TestDelegateSplitOtherTokens(t *testing.T) {
	lang, _ := makeDelegationTestLexers()
	it, err := lang.Tokenise(nil, delegateSourceMiddle)
	assert.NoError(t, err)
	splits, other := splitOtherTokens(it)
	assert.Equal(t, "hello world  there", other)
	expected := []tokenSplit{tokenSplit{
		pos: 12,
		tokens: []*Token{
			{CommentPreproc, "<?"},
			{Whitespace, " "},
			{Keyword, "what"},
			{Whitespace, " "},
			{CommentPreproc, "?>"},
		},
	}}
	assert.Equal(t, expected, splits)
}

func TestDelegateSplitOtherTokensSourceAtEnd(t *testing.T) {
	lang, _ := makeDelegationTestLexers()
	lang = Coalesce(lang)
	it, err := lang.Tokenise(nil, delegateSourceEnd)
	assert.NoError(t, err)
	splits, other := splitOtherTokens(it)
	assert.Equal(t, "hello world ", other)
	expected := []tokenSplit{tokenSplit{
		pos: 12,
		tokens: []*Token{
			{CommentPreproc, "<?"},
			{Whitespace, " "},
			{Keyword, "what"},
			{TextWhitespace, " "},
			{Error, "there"},
		},
	}}
	assert.Equal(t, expected, splits)
}

func TestDelegate(t *testing.T) {
	lang, root := makeDelegationTestLexers()
	delegate := DelegatingLexer(root, lang)
	it, err := delegate.Tokenise(nil, delegateSourceMiddle)
	assert.NoError(t, err)
	expected := []*Token{
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
	}
	assert.Equal(t, expected, it.Tokens())
}

func TestDelegateEnd(t *testing.T) {
	lang, root := makeDelegationTestLexers()
	lang = Coalesce(lang)
	delegate := DelegatingLexer(root, lang)
	it, err := delegate.Tokenise(nil, delegateSourceEnd)
	assert.NoError(t, err)
	expected := []*Token{
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
	}
	assert.Equal(t, expected, it.Tokens())
}
