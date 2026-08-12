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
