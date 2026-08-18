package styles

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v3"
)

// TestEmbeddedStyleNamesMatchFilenames verifies that every embedded style
// parses and is stored in a file named after its lowercased style name. The
// lazy filename-based lookup in this package depends on this invariant.
func TestEmbeddedStyleNamesMatchFilenames(t *testing.T) {
	names := embeddedNames()
	assert.True(t, len(names) > 0)
	for _, name := range names {
		f, err := embedded.Open(name + ".xml")
		assert.NoError(t, err)
		style, err := chroma.NewXMLStyle(f)
		_ = f.Close()
		assert.NoError(t, err)
		assert.Equal(t, name, strings.ToLower(style.Name),
			"%s.xml: file name must be the lowercased style name", name)
	}
}
