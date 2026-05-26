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

func TestGetForMode(t *testing.T) {
	tests := []struct {
		name string
		req  string
		mode chroma.Mode
		want string
	}{
		{"AlreadyDark", "github-dark", chroma.Dark, "github-dark"},
		{"AlreadyLight", "github", chroma.Light, "github"},
		{"FollowsCounterpartToDark", "github", chroma.Dark, "github-dark"},
		{"FollowsCounterpartToLight", "github-dark", chroma.Light, "github"},
		{"NoCounterpartReturnsOriginal", "swapoff", chroma.Light, "swapoff"},
		{"UnknownReturnsFallback", "no-such-style", chroma.Dark, Fallback.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetForMode(tt.req, tt.mode)
			assert.Equal(t, tt.want, got.Name)
		})
	}
}

func TestRegisterPair(t *testing.T) {
	light := chroma.MustNewStyle("PairLight", chroma.StyleEntries{
		chroma.Background: "bg:#ffffff",
	})
	dark := chroma.MustNewStyle("PairDark", chroma.StyleEntries{
		chroma.Background: "bg:#000000",
	})
	RegisterPair(light, dark)
	assert.Equal(t, "pairdark", light.Counterpart)
	assert.Equal(t, "pairlight", dark.Counterpart)
	assert.Equal(t, dark, GetForMode("PairLight", chroma.Dark))
	assert.Equal(t, light, GetForMode("PairDark", chroma.Light))
}
