package formatters

import (
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v2"
)

func TestClosestColour(t *testing.T) {
	actual := findClosest(ttyTables[256], chroma.MustParseColour("#e06c75"))
	assert.Equal(t, chroma.MustParseColour("#d75f87"), actual)
}

func TestNoneColour(t *testing.T) {
	formatter := TTY256
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

	// "178" = #d7af00 approximates #d0ab1e
	//
	// 178 color ref: https://jonasjacek.github.io/colors/
	assert.Equal(t, "\033[38;5;178mWORD\033[0m", stringBuilder.String())
}
