package lexers

import (
	"io/fs"
	"testing"

	assert "github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v3"
)

// TestEmbeddedLexerMetadata verifies that the generated metadata in
// lexers_gen.go matches the <config> element of every embedded XML
// definition. If this fails, run "go generate ./lexers".
func TestEmbeddedLexerMetadata(t *testing.T) {
	paths, err := fs.Glob(embedded, "embedded/*.xml")
	assert.NoError(t, err)
	assert.Equal(t, len(paths), len(embeddedLexers))
	for i, path := range paths {
		entry := &embeddedLexers[i]
		assert.Equal(t, path, entry.path)
		lexer, err := chroma.NewXMLLexer(embedded, path)
		assert.NoError(t, err)
		assert.Equal(t, *lexer.Config(), entry.config, "%s: generated metadata is stale, run 'go generate ./lexers'", path)
	}
}
