package lexers_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
)

func TestYAMLNumericPrefixWithLongWhitespace(t *testing.T) {
	spaces := strings.Repeat(" ", 10000)
	for _, test := range []struct {
		name, value, literal, comment string
	}{
		{"PlainSuffix", "2" + spaces + "suffix", "2" + spaces + "suffix", ""},
		{"AttachedHash", "2" + spaces + "suffix#literal", "2" + spaces + "suffix#literal", ""},
		{"CommentAfterSuffix", "2" + spaces + "suffix # comment", "2" + spaces + "suffix", "# comment"},
		{"SpacesInComment", "2 suffix # comment" + spaces + "tail", "2 suffix", "# comment" + spaces + "tail"},
	} {
		t.Run(test.name, func(t *testing.T) {
			it, err := lexers.Get("yaml").Tokenise(nil, "v: "+test.value+"\n")
			assert.NoError(t, err)
			expected := []chroma.Token{
				{Type: chroma.NameTag, Value: "v"},
				{Type: chroma.Punctuation, Value: ":"},
				{Type: chroma.TextWhitespace, Value: " "},
				{Type: chroma.Literal, Value: test.literal},
			}
			if test.comment != "" {
				expected = append(expected,
					chroma.Token{Type: chroma.TextWhitespace, Value: " "},
					chroma.Token{Type: chroma.Comment, Value: test.comment})
			}
			expected = append(expected, chroma.Token{Type: chroma.TextWhitespace, Value: "\n"})
			assert.Equal(t, expected, slices.Collect(it))
		})
	}
}
