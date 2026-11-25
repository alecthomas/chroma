package html

import (
	"io"
	"iter"
	"testing"

	assert "github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/styles"
)

func TestFormatRecoversIteratorPanic(t *testing.T) {
	panicking := func(func(chroma.Token) bool) { panic("lexer failed") }
	err := New().Format(io.Discard, styles.Fallback, iter.Seq[chroma.Token](panicking))
	assert.Error(t, err)
	assert.Equal(t, "lexer failed", err.Error())
}
