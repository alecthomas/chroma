package svg

import (
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v3"
)

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

func TestTokenWidthIsMeasuredInColumns(t *testing.T) {
	style := chroma.MustNewStyle("svg-test", chroma.StyleEntries{
		chroma.Background: "#000000 bg:#ffffff",
		chroma.Comment:    "#008000 bg:#eeeeee",
	})
	render := func(text string) string {
		w := &strings.Builder{}
		assert.NoError(t, New().Format(w, style, chroma.Literator(
			chroma.Token{Type: chroma.Text, Value: text},
			chroma.Token{Type: chroma.Comment, Value: "// x"},
		)))
		return w.String()
	}

	// "héllo" is five columns wide but six bytes long.
	t.Run("Canvas", func(t *testing.T) {
		assert.Contains(t, render("héllo"), `<svg width="72px"`)
	})
	t.Run("Background", func(t *testing.T) {
		assert.Contains(t, render("héllo"), `x="5ch" y="0.250000em" width="4ch"`)
	})
	// "ＡＢ" is fullwidth: four columns wide and six bytes long.
	t.Run("WideCanvas", func(t *testing.T) {
		assert.Contains(t, render("ＡＢ"), `<svg width="64px"`)
	})
	t.Run("WideBackground", func(t *testing.T) {
		assert.Contains(t, render("ＡＢ"), `x="4ch" y="0.250000em" width="4ch"`)
	})
}
