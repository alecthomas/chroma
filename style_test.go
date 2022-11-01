package chroma

import (
	"encoding/xml"
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestStyleInherit(t *testing.T) {
	s, err := NewStyle("test", StyleEntries{
		Name:         "bold #f00",
		NameVariable: "#fff",
	})
	assert.NoError(t, err)
	assert.Equal(t, StyleEntry{Colour: 0x1000000, Bold: Yes}, s.Get(NameVariable))
}

func TestStyleColours(t *testing.T) {
	s, err := NewStyle("test", StyleEntries{
		Name: "#f00 bg:#001 border:#ansiblue",
	})
	assert.NoError(t, err)
	assert.Equal(t, StyleEntry{Colour: 0xff0001, Background: 0x000012, Border: 0x000100}, s.Get(Name))
}

func TestStyleClone(t *testing.T) {
	parent, err := NewStyle("test", StyleEntries{
		Background: "bg:#ffffff",
	})
	assert.NoError(t, err)
	clone, err := parent.Builder().Add(Comment, "#0f0").Build()
	assert.NoError(t, err)

	assert.Equal(t, "bg:#ffffff", clone.Get(Background).String())
	assert.Equal(t, "#00ff00 bg:#ffffff", clone.Get(Comment).String())
	assert.Equal(t, "bg:#ffffff", parent.Get(Comment).String())
}

func TestSynthesisedStyleEntries(t *testing.T) {
	style, err := NewStyle("test", StyleEntries{
		Background: "bg:#ffffff",
	})
	assert.NoError(t, err)
	assert.True(t, style.Has(LineHighlight))
	assert.True(t, style.Has(LineNumbersTable))
	assert.True(t, style.Has(LineNumbers))
	assert.Equal(t, "bg:#e5e5e5", style.Get(LineHighlight).String())
	assert.Equal(t, "#7f7f7f bg:#ffffff", style.Get(LineNumbers).String())
	assert.Equal(t, "#7f7f7f bg:#ffffff", style.Get(LineNumbersTable).String())
}

func TestSynthesisedStyleClone(t *testing.T) {
	style, err := NewStyle("test", StyleEntries{
		Background:    "bg:#ffffff",
		LineHighlight: "bg:#ffffff",
		LineNumbers:   "bg:#fffff1",
	})
	assert.NoError(t, err)
	style, err = style.Builder().Build()
	assert.NoError(t, err)
	assert.True(t, style.Has(LineHighlight))
	assert.True(t, style.Has(LineNumbers))
	assert.Equal(t, "bg:#ffffff", style.Get(LineHighlight).String())
	assert.Equal(t, "bg:#fffff1", style.Get(LineNumbers).String())
}

func TestStyleBuilderTransform(t *testing.T) {
	orig, err := NewStyle("test", StyleEntries{
		Name:         "#000",
		NameVariable: "bold #f00",
	})
	assert.NoError(t, err)

	// Derive a style that inherits entries from orig.
	builder := orig.Builder()
	builder.Add(NameVariableGlobal, "#f30")
	deriv, err := builder.Build()
	assert.NoError(t, err)

	// Use Transform to brighten or darken all of the colours in the derived style.
	light, err := deriv.Builder().Transform(func(se StyleEntry) StyleEntry {
		se.Colour = se.Colour.ClampBrightness(0.9, 1)
		return se
	}).Build()
	assert.NoError(t, err, "Transform failed: %v", err)
	assert.True(t, light.Get(Name).Colour.Brightness() >= 0.89)
	assert.True(t, light.Get(NameVariable).Colour.Brightness() >= 0.89)
	assert.True(t, light.Get(NameVariableGlobal).Colour.Brightness() >= 0.89)

	dark, err := deriv.Builder().Transform(func(se StyleEntry) StyleEntry {
		se.Colour = se.Colour.ClampBrightness(0, 0.1)
		return se
	}).Build()
	assert.NoError(t, err, "Transform failed: %v", err)
	assert.True(t, dark.Get(Name).Colour.Brightness() <= 0.11)
	assert.True(t, dark.Get(NameVariable).Colour.Brightness() <= 0.11)
	assert.True(t, dark.Get(NameVariableGlobal).Colour.Brightness() <= 0.11)

	// The original styles should be unchanged.
	assert.Equal(t, "#000000", orig.Get(Name).Colour.String())
	assert.Equal(t, "#ff0000", orig.Get(NameVariable).Colour.String())
	assert.Equal(t, "#ff3300", deriv.Get(NameVariableGlobal).Colour.String())
}

func TestStyleMarshaller(t *testing.T) {
	expected, err := NewStyle("test", StyleEntries{
		Whitespace: "bg:#ffffff",
		Text:       "#000000 underline",
	})
	assert.NoError(t, err)
	data, err := xml.MarshalIndent(expected, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, `<style name="test">
  <entry type="Text" style="underline #000000"></entry>
  <entry type="TextWhitespace" style="bg:#ffffff"></entry>
</style>`, string(data))
	actual := &Style{}
	err = xml.Unmarshal(data, actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
