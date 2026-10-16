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
