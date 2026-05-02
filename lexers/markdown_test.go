package lexers

import (
	"fmt"
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v2"
)

func TestSplitFrontmatter(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		frontmatter string
		rest        string
		ok          bool
	}{
		{
			name:        "leading block",
			input:       "---\nkey: value\n---\nbody\n",
			frontmatter: "---\nkey: value\n---\n",
			rest:        "body\n",
			ok:          true,
		},
		{
			name:  "horizontal rule only",
			input: "---\nbody\n",
			rest:  "---\nbody\n",
			ok:    false,
		},
		{
			name:  "mid document block",
			input: "body\n---\nkey: value\n---\n",
			rest:  "body\n---\nkey: value\n---\n",
			ok:    false,
		},
		{
			name:  "unterminated block",
			input: "---\nkey: value\nbody\n",
			rest:  "---\nkey: value\nbody\n",
			ok:    false,
		},
		{
			name:        "crlf block",
			input:       "---\r\nkey: value\r\n---\r\nbody\r\n",
			frontmatter: "---\r\nkey: value\r\n---\r\n",
			rest:        "body\r\n",
			ok:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			frontmatter, rest, ok := splitFrontmatter(test.input)
			assert.Equal(t, test.ok, ok)
			assert.Equal(t, test.frontmatter, frontmatter)
			assert.Equal(t, test.rest, rest)
		})
	}
}

func TestMarkdownFrontmatterAndCommentStress(t *testing.T) {
	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("title: stress test\n")
	b.WriteString("summary: markdown frontmatter and comment parsing\n")
	b.WriteString("layout: docs\n")
	b.WriteString("---\n\n")

	for i := range 16 {
		fmt.Fprintf(&b, "Paragraph %03d <!-- comment %03d with --- and title: not frontmatter -->\n", i, i)
		if i%8 == 0 {
			fmt.Fprintf(&b, "---\nbody_key_%03d: should stay text\n---\n", i)
		}
		if i%16 == 0 {
			b.WriteString("---\n\n")
		}
		b.WriteString("\n")
	}

	input := b.String()
	tokens, err := chroma.Tokenise(Markdown, nil, input)
	assert.NoError(t, err)
	assert.Equal(t, input, chroma.Stringify(tokens...))

	var commentCount int
	var frontmatterNameTags int
	var bodyNameTags int
	for _, token := range tokens {
		assert.NotEqual(t, chroma.Error, token.Type)

		switch {
		case token.Type == chroma.CommentMultiline && strings.HasPrefix(token.Value, "<!-- comment "):
			commentCount++
		case token.Type == chroma.NameTag && (token.Value == "title" || token.Value == "summary" || token.Value == "layout"):
			frontmatterNameTags++
		case token.Type == chroma.NameTag && strings.HasPrefix(token.Value, "body_key_"):
			bodyNameTags++
		}
	}

	assert.Equal(t, 16, commentCount)
	assert.Equal(t, 3, frontmatterNameTags)
	assert.Zero(t, bodyNameTags)
}
