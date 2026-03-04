package styles

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v2"
)

func TestStyleRegistryCaseInsensitivity(t *testing.T) {
	// Verify that all keys in the Registry are lowercase.
	for name := range Registry {
		assert.Equal(t, strings.ToLower(name), name)
	}

	// Verify that Get is case-insensitive.
	names := []string{"monokai", "Monokai", "MONOKAI"}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			style := Get(name)

			assert.NotEqual(t, Fallback, style)
			assert.True(t, strings.EqualFold(style.Name, name))
		})
	}
}

func TestGetUnknownStyleReturnsFallback(t *testing.T) {
	assert.Equal(t, Fallback, Get("non-existent-style"))
}

func TestRegisterCaseInsensitivity(t *testing.T) {
	custom := chroma.MustNewStyle("CustomStyle", chroma.StyleEntries{
		chroma.Text: "#ffffff",
	})
	Register(custom)

	assert.Equal(t, custom, Get("customstyle"))
	assert.Equal(t, custom, Get("CUSTOMSTYLE"))
	assert.Equal(t, custom, Get("CustomStyle"))
}
