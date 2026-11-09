package lexers_test

import (
	"slices"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
)

func TestJSONComments(t *testing.T) {
	for _, source := range []string{
		"// comment at EOF",
		"{\"value\": 1} // comment at EOF",
		"/* block\ncomment */ {\"value\": 1}",
		"{ /* before key */ \"value\" /* before colon */ : /* before value */ 1 /* after value */ }",
		"[1, /* between items */ 2]",
	} {
		t.Run(source, func(t *testing.T) {
			lexer := lexers.Match("settings.jsonc")
			it, err := chroma.Tokenise(lexer, nil, source)
			assert.NoError(t, err)
			tokens := slices.Collect(it)
			assert.Equal(t, source, chroma.Stringify(tokens...))
			for _, token := range tokens {
				assert.NotEqual(t, chroma.Error, token.Type)
			}
		})
	}
}
