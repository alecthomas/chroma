package lexers_test

import (
	"slices"
	"testing"
	"time"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
)

func TestJungleLexerDoesNotHangOnUnhandledPunctuation(t *testing.T) {
	lx := lexers.Get("Jungle")
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
			case <-done:
			case <-time.After(2 * time.Second):
				t.Fatalf("Tokenise(%q) never returned", in)
			}
		})
	}
}
