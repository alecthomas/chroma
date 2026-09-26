package svg

import (
	"strings"
	"testing"

	"github.com/alecthomas/chroma/v3"

	assert "github.com/alecthomas/assert/v2"
)

func TestNonASCIIWidths(t *testing.T) {
	style, err := chroma.NewStyle("test", chroma.StyleEntries{
		chroma.Background:    "#ffffff bg:#ffffff",
		chroma.LiteralString: "bg:#fff0f0",
	})
	assert.NoError(t, err)

	tokens := []chroma.Token{
		{Type: chroma.Text, Value: "x = "},
		{Type: chroma.LiteralString, Value: `"éé"`},
		{Type: chroma.Text, Value: ", "},
		{Type: chroma.LiteralString, Value: `"ab"`},
		{Type: chroma.Text, Value: "\n"},
	}

	w := &strings.Builder{}
	assert.NoError(t, New().Format(w, style, chroma.Literator(tokens...)))
	out := w.String()

	// Widths and offsets are in character cells, so multi-byte runes count once.
	assert.Contains(t, out, `<svg width="120px"`)
	assert.Contains(t, out, `x="4ch" y="0.250000em" width="4ch"`)
	assert.Contains(t, out, `x="10ch" y="0.250000em" width="4ch"`)
}

func TestWriteFontStyle(t *testing.T) {
	tests := []struct {
		name     string
		format   FontFormat
		expected string
	}{
		{"WOFF", WOFF, "src: url(data:font/woff;charset=utf-8;base64,AAAA) format('woff');\n"},
		{"WOFF2", WOFF2, "src: url(data:font/woff2;charset=utf-8;base64,AAAA) format('woff2');\n"},
		{"TrueType", TRUETYPE, "src: url(data:font/ttf;charset=utf-8;base64,AAAA) format('truetype');\n"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := New(EmbedFont("Test", "AAAA", test.format))
			w := &strings.Builder{}
			assert.NoError(t, f.writeFontStyle(w))
			assert.Contains(t, w.String(), test.expected)
		})
	}
}
