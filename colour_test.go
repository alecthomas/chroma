package chroma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColourRGB(t *testing.T) {
	colour := ParseColour("#8913af")
	require.Equal(t, uint8(0x89), colour.Red())
	require.Equal(t, uint8(0x13), colour.Green())
	require.Equal(t, uint8(0xaf), colour.Blue())
}

func TestColourString(t *testing.T) {
	require.Equal(t, "#8913af", ParseColour("#8913af").String())
}
