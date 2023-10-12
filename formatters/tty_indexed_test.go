package formatters

import (
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
)

func TestClosestColour(t *testing.T) {
	actual := findClosest(ttyTables[256], chroma.MustParseColour("#e06c75"))
	assert.Equal(t, chroma.MustParseColour("#d75f87"), actual)
}

func TestNoneColour(t *testing.T) {
	style := styles.Registry["gruvbox"]
	formatter := TTY256
	tokenType := chroma.None

	stringBuilder := strings.Builder{}
	err := formatter.Format(&stringBuilder, style, chroma.Literator(chroma.Token{
		Type:  tokenType,
		Value: "WORD",
	}))
	assert.NoError(t, err)

	// "187" = #d7d7af approximates #ebdbb2, which is the Gruvbox foreground
	// color of the "Background" type, see gruvbox.xml in this repo.
	//
	// 187 color ref: https://jonasjacek.github.io/colors/
	assert.Equal(t, "\033[38;5;187mWORD\033[0m", stringBuilder.String())
}
