package lexers_test

import (
	"testing"
	"time"

	"github.com/alecthomas/chroma/v2/lexers"
)

// TestJungleHangsOnUnhandledPunctuation reproduces a hang in the bundled
// "Jungle" lexer: "instruction" pushes "var" on `(?=\S)`, and "var"'s
// catch-all rule pops back on any character it doesn't special-case (root's
// [\.;\[\]\(\)\$] and \w+ are the only things "var" consumes). Root cause:
// see TestZeroWidthPushPopCausesInfiniteLoop in the chroma package.
func TestJungleHangsOnUnhandledPunctuation(t *testing.T) {
	lx := lexers.Get("Jungle")
	if lx == nil {
		t.Fatal("no Jungle lexer registered")
	}
	for _, in := range []string{"/", "*", ":", `"`, "!", "%", "'"} {
		in := in
		done := make(chan struct{})
		go func() {
			defer close(done)
			it, err := lx.Tokenise(nil, in)
			if err != nil {
				t.Errorf("Tokenise(%q) error: %v", in, err)
				return
			}
			_ = it.Tokens()
		}()
		select {
		case <-done:
		case <-time.After(1 * time.Second):
			t.Errorf("Tokenise(%q) never returned: instruction<->var loop never consumes the character", in)
		}
	}
}
