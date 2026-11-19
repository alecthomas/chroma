package formatters

import (
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v3"
)

func TestClosestIRCColour(t *testing.T) {
	actual := findClosestIRCColor("#6c329c")
	assert.Equal(t, "49", actual)
}

func TestNoneIRCColour(t *testing.T) {
	formatter := MIRC100
	tokenType := chroma.None

	style, err := chroma.NewStyle("test", chroma.StyleEntries{
		chroma.Background: "#D0ab1e",
	})
	assert.NoError(t, err)

	stringBuilder := strings.Builder{}
	err = formatter.Format(&stringBuilder, style, chroma.Literator(chroma.Token{
		Type:  tokenType,
		Value: "WORD",
	}))
	assert.NoError(t, err)

	assert.Equal(t, "\x0342WORD\x0f", stringBuilder.String())
}

func TestHexIRCBackgroundColour(t *testing.T) {
	style, err := chroma.NewStyle("test", chroma.StyleEntries{
		chroma.Error: "#960050 bg:#1e0010",
	})
	assert.NoError(t, err)

	stringBuilder := strings.Builder{}
	err = MIRC16m.Format(&stringBuilder, style, chroma.Literator(chroma.Token{
		Type:  chroma.Error,
		Value: "WORD",
	}))
	assert.NoError(t, err)
	assert.Equal(t, "\x04960050,1E0010WORD\x0f", stringBuilder.String())
}

func TestIRCLeadingComma(t *testing.T) {
	for _, tc := range []struct {
		name, style, text, want string
		formatter               chroma.Formatter
	}{
		{"indexed", "#ff0000", ",13,units", "\x0304\x02\x02,13,units\x0f", MIRC100},
		{"hex multiline", "bold #ff0000", "first\r\n,aBcDeF", "\x02\x04FF0000first\x0f\r\n\x02\x04FF0000\x02\x02,aBcDeF\x0f", MIRC16m},
		{"indexed background", "#ff0000 bg:#000000", ",13", "\x0304,01,13\x0f", MIRC100},
		{"short hex", "#ff0000", ",12345", "\x04FF0000,12345\x0f", MIRC16m},
	} {
		t.Run(tc.name, func(t *testing.T) {
			style := chroma.MustNewStyle("test", chroma.StyleEntries{chroma.LiteralString: tc.style})
			var out strings.Builder
			err := tc.formatter.Format(&out, style, chroma.Literator(chroma.Token{Type: chroma.LiteralString, Value: tc.text}))
			assert.NoError(t, err)
			assert.Equal(t, tc.want, out.String())
		})
	}
}
