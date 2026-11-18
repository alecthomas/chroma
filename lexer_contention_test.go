package chroma_test

import (
	"bytes"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/lexers"
	"github.com/alecthomas/chroma/v3/styles"
)

// Regression test for #1498: under CPU contention a rule match can spuriously
// exceed its 250ms timeout, which must not fail tokenising.
func TestTokeniseUnderCPUContention(t *testing.T) {
	old := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(old)
	var stop atomic.Bool
	for range 64 {
		go func() {
			for !stop.Load() {
			}
		}()
	}
	defer stop.Store(true)
	tokens, err := lexers.Get("go").Tokenise(nil, strings.Repeat("s := \"a\" + \"b\"\n", 1000))
	assert.NoError(t, err)
	var out bytes.Buffer
	assert.NoError(t, html.New().Format(&out, styles.Fallback, tokens))
	assert.True(t, out.Len() > 0)
}
