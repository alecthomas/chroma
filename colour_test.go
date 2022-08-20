package chroma

import (
	"math"
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

// hue returns c's hue. See https://stackoverflow.com/a/23094494.
func hue(c Colour) float64 {
	r := float64(c.Red()) / 255
	g := float64(c.Green()) / 255
	b := float64(c.Blue()) / 255

	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)

	switch {
	case r == min:
		return (g - b) / (max - min)
	case g == min:
		return 2 + (b-r)/(max-min)
	default:
		return 4 + (r-g)/(max-min)
	}
}

func TestColourClampBrightness(t *testing.T) {
	const delta = 0.01 // used for brightness and hue comparisons

	// Start with a colour with a brightness close to 0.5.
	initial := NewColour(0, 128, 255)
	br := initial.Brightness()
	assert.InDelta(t, 0.5, br, delta)

	// Passing a range that includes the colour's brightness should be a no-op.
	assert.Equal(t, initial.String(), initial.ClampBrightness(br-0.01, br+0.01).String())

	// Clamping to [0, 0] or [1, 1] should produce black or white, respectively.
	assert.Equal(t, "#000000", initial.ClampBrightness(0, 0).String())
	assert.Equal(t, "#ffffff", initial.ClampBrightness(1, 1).String())

	// Clamping to a brighter or darker range should produce the requested
	// brightness while preserving the colour's hue.
	brighter := initial.ClampBrightness(0.75, 1)
	assert.InDelta(t, 0.75, brighter.Brightness(), delta)
	assert.InDelta(t, hue(initial), hue(brighter), delta)

	darker := initial.ClampBrightness(0, 0.25)
	assert.InDelta(t, 0.25, darker.Brightness(), delta)
	assert.InDelta(t, hue(initial), hue(darker), delta)
}
