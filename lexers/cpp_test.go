package lexers_test

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
)

func TestIssue290(t *testing.T) {
	input := `// 64-bit floats have 53 digits of precision, including the whole-number-part.
double a =     0011111110111001100110011001100110011001100110011001100110011010; // imperfect representation of 0.1
double b =     0011111111001001100110011001100110011001100110011001100110011010; // imperfect representation of 0.2
double c =     0011111111010011001100110011001100110011001100110011001100110011; // imperfect representation of 0.3
double a + b = 0011111111010011001100110011001100110011001100110011001100110100; // Note that this is not quite equal to the "canonical" 0.3!a
`
	it, err := lexers.GlobalLexerRegistry.Get("C++").Tokenise(nil, input)
	assert.NoError(t, err)
	for {
		token := it()
		if token == chroma.EOF {
			break
		}
	}
}
