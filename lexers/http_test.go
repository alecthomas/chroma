package lexers

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v2"
)

// Regression test for https://github.com/alecthomas/chroma/issues/1273:
// the HTTP lexer used to drop body tokens when not wrapped in chroma.Coalesce.
func TestHTTPBodyTokensWithoutCoalesce(t *testing.T) {
	source := `GET /foo HTTP/1.1
Content-Type: application/json
User-Agent: foo

{"hello": "world"}
`
	tokens, err := chroma.Tokenise(HTTP, nil, source)
	assert.NoError(t, err)
	assert.Equal(t, source, chroma.Stringify(tokens...))

	found := false
	for _, token := range tokens {
		if token.Type == chroma.LiteralStringDouble && token.Value == `"world"` {
			found = true
		}
	}
	assert.True(t, found, "expected body to be tokenised by sub-lexer")
}
