package chroma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemappingLexer(t *testing.T) {
	var lexer Lexer = MustNewLexer(nil, Rules{ // nolint: forbidigo
		"root": {
			{`\s+`, Whitespace, nil},
			{`\w+`, Name, nil},
		},
	})
	lexer = TypeRemappingLexer(lexer, TypeMapping{
		{Name, Keyword, []string{"if", "else"}},
	})

	it, err := lexer.Tokenise(nil, `if true then print else end`)
	assert.NoError(t, err)
	expected := []Token{
		{Keyword, "if"}, {TextWhitespace, " "}, {Name, "true"}, {TextWhitespace, " "}, {Name, "then"},
		{TextWhitespace, " "}, {Name, "print"}, {TextWhitespace, " "}, {Keyword, "else"},
		{TextWhitespace, " "}, {Name, "end"},
	}
	actual := it.Tokens()
	assert.Equal(t, expected, actual)
}
