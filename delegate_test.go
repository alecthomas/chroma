package chroma

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func makeDelegationTestLexers(t *testing.T) (lang Lexer, root Lexer) {
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
		// nolint: scopelint
		t.Run(test.name, func(t *testing.T) {
			it, err := delegate.Tokenise(nil, test.source)
			assert.NoError(t, err)
			actual := it.Tokens()
			assert.Equal(t, test.expected, actual)
		})
	}
}
