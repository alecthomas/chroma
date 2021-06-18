package chroma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColourRGB(t *testing.T) {
	colour := ParseColour("#8913af")
	assert.Equal(t, uint8(0x89), colour.Red())
	assert.Equal(t, uint8(0x13), colour.Green())
	assert.Equal(t, uint8(0xaf), colour.Blue())
}

func TestColourString(t *testing.T) {
	assert.Equal(t, "#8913af", ParseColour("#8913af").String())
}

func distance(a, b uint8) uint8 {
	if a < b {
		return b - a
	}
	return a - b
}

func TestColourBrighten(t *testing.T) {
	actual := NewColour(128, 128, 128).Brighten(0.5)
	// Closeish to what we expect is fine.
	assert.True(t, distance(192, actual.Red()) <= 2)
	assert.True(t, distance(192, actual.Blue()) <= 2)
	assert.True(t, distance(192, actual.Green()) <= 2)
	actual = NewColour(128, 128, 128).Brighten(-0.5)
	assert.True(t, distance(65, actual.Red()) <= 2)
	assert.True(t, distance(65, actual.Blue()) <= 2)
	assert.True(t, distance(65, actual.Green()) <= 2)
}

func TestColourBrightess(t *testing.T) {
	actual := NewColour(128, 128, 128).Brightness()
	assert.True(t, distance(128, uint8(actual*255.0)) <= 2)
}
