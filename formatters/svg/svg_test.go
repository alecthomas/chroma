package svg

import (
	"encoding/xml"
	"io"
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/styles"
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

func TestFontFamilyIsEscaped(t *testing.T) {
	for _, fontFamily := range []string{`"Fira Code", monospace`, `Foo & Bar Mono`} {
		t.Run(fontFamily, func(t *testing.T) {
			w := &strings.Builder{}
			f := New(EmbedFont(fontFamily, "AAAA", WOFF))
			assert.NoError(t, f.Format(w, styles.Get("monokai"),
				chroma.Literator(chroma.Token{Type: chroma.Text, Value: "hi\n"})))
			attr, style := decodeSVG(t, w.String())
			assert.Equal(t, fontFamily, attr)
			assert.Contains(t, style, "font-family: '"+fontFamily+"';")
		})
	}
}

func decodeSVG(t *testing.T, doc string) (fontFamily, style string) {
	t.Helper()
	dec := xml.NewDecoder(strings.NewReader(doc))
	dec.Strict = true
	inStyle := false
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return fontFamily, style
		}
		assert.NoError(t, err, "SVG output is not well-formed XML")
		switch tok := tok.(type) {
		case xml.StartElement:
			inStyle = tok.Name.Local == "style"
			if tok.Name.Local == "g" {
				for _, a := range tok.Attr {
					if a.Name.Local == "font-family" {
						fontFamily = a.Value
					}
				}
			}
		case xml.CharData:
			if inStyle {
				style += string(tok)
			}
		case xml.EndElement:
			inStyle = false
		}
	}
}
