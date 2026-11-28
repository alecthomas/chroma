package formatters

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestLookup(t *testing.T) {
	f, ok := Lookup("json")
	assert.True(t, ok)
	assert.Equal(t, JSON, f)

	f, ok = Lookup("missing")
	assert.False(t, ok)
	assert.Zero(t, f)
	assert.Equal(t, Fallback, Get("missing"))
}
