package lexers

import (
	"slices"
	"testing"
	"time"

	"github.com/alecthomas/chroma/v3"
)

func TestJungleLexerDoesNotHangOnUnhandledPunctuation(t *testing.T) {
	lx := Get("Jungle")
	if lx == nil {
		t.Fatal("no Jungle lexer registered")
	}
	for _, in := range []string{"/", "*", ":", `"`, "!", "%", "'"} {
		t.Run(in, func(t *testing.T) {
			it, err := lx.Tokenise(nil, in)
			if err != nil {
				t.Fatalf("Tokenise(%q): %v", in, err)
			}
			done := make(chan []chroma.Token, 1)
			go func() {
				done <- slices.Collect(it)
			}()
			select {
			case tokens := <-done:
				if len(tokens) == 0 {
					t.Fatalf("expected tokens for %q, got none", in)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("Tokenise(%q) never returned", in)
			}
		})
	}
}
