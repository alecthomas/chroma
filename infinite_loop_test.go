package chroma

import (
	"testing"
	"time"
)

// TestZeroWidthPushPopCausesInfiniteLoop shows LexerState.Iterator has no
// guard against a push/pop cycle that never advances Pos: "a" pushes "b" on
// a zero-width lookahead, "b" unconditionally pops back to "a", and neither
// ever matches the character itself. Real-world trigger: the bundled
// "Jungle" lexer (see TestJungleHangsOnUnhandledPunctuation in lexers/).
func TestZeroWidthPushPopCausesInfiniteLoop(t *testing.T) {
	lexer := mustNewLexer(t, &Config{Name: "loopy"}, Rules{
		"root": {
			{`(?=\S)`, None, Push("a")},
		},
		"a": {
			{`(?=\S)`, None, Push("b")},
		},
		"b": {
			{``, None, Pop(1)},
		},
	})

	done := make(chan struct{})
	go func() {
		defer close(done)
		it, err := lexer.Tokenise(nil, "x")
		if err != nil {
			t.Errorf("Tokenise error: %v", err)
			return
		}
		_ = it.Tokens()
	}()

	select {
	case <-done:
		t.Log("returned fine")
	case <-time.After(2 * time.Second):
		t.Fatal("Iterator never returned: zero-width push/pop cycle spins forever without consuming input")
	}
}
