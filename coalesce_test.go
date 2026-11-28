package chroma

import (
	"errors"
	"iter"
	"slices"
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestCoalesce(t *testing.T) {
	lexer := Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": []Rule{
			{`[!@#$%^&*()]`, Punctuation, nil},
		},
	}))
	actual, err := lexer.Tokenise("!@#$")
	assert.NoError(t, err)
	expected := []Token{{Punctuation, "!@#$"}}
	assert.Equal(t, expected, slices.Collect(actual))
}

func TestCoalesceFluentSetters(t *testing.T) {
	for _, method := range []string{"analyser", "registry", "both"} {
		t.Run(method, func(t *testing.T) {
			base := mustNewLexer(t, &Config{Name: "coalesce-setters"}, Rules{"root": {{`[!@#$]`, Punctuation, nil}}})
			original := Coalesce(base)
			if original.Config() != base.Config() {
				t.Error("config not forwarded")
			}
			configured := original
			registry := &LexerRegistry{}
			if method == "analyser" || method == "both" {
				configured = configured.SetAnalyser(func(string) float32 { return 0.75 })
			}
			if method == "registry" || method == "both" {
				configured = configured.SetRegistry(registry)
			}
			if configured != original {
				t.Errorf("setter returned %T, want original coalescer %T", configured, original)
			}
			if method != "registry" && configured.AnalyseText("anything") != 0.75 {
				t.Error("analyser not forwarded")
			}
			if method != "analyser" && base.registry != registry {
				t.Error("registry not forwarded")
			}
			seq, err := configured.Tokenise("!@#$")
			if err != nil {
				t.Fatal(err)
			}
			got := slices.Collect(seq)
			if len(got) != 1 || got[0] != (Token{Punctuation, "!@#$"}) {
				t.Errorf("tokens=%v, want one coalesced token", got)
			}
		})
	}
}

type coalesceSetterLexer struct {
	Lexer
	tokenise func(string, ...TokeniseOption) (iter.Seq[Token], error)
}

func (l *coalesceSetterLexer) Tokenise(text string, options ...TokeniseOption) (iter.Seq[Token], error) {
	return l.tokenise(text, options...)
}

func TestCoalesceConfiguredErrorAndEarlyStop(t *testing.T) {
	base := mustNewLexer(t, &Config{Name: "setter-controls"}, Rules{"root": {{`.`, Text, nil}}})
	failure := errors.New("tokenise failed")
	wrapped := &coalesceSetterLexer{Lexer: base, tokenise: func(string, ...TokeniseOption) (iter.Seq[Token], error) { return nil, failure }}
	configured := Coalesce(wrapped).SetAnalyser(func(string) float32 { return 0.5 }).SetRegistry(&LexerRegistry{})
	if seq, err := configured.Tokenise("x"); seq != nil || err != failure {
		t.Fatalf("Tokenise()=%v,%v; want nil, original error", seq, err)
	}
	calls, closed := 0, false
	wrapped.tokenise = func(string, ...TokeniseOption) (iter.Seq[Token], error) {
		return func(yield func(Token) bool) {
			defer func() { closed = true }()
			for _, token := range []Token{{Text, "a"}, {Text, "b"}, {Punctuation, "!"}, {Text, "c"}} {
				calls++
				if !yield(token) {
					return
				}
			}
		}, nil
	}
	seq, err := configured.Tokenise("")
	if err != nil {
		t.Fatal(err)
	}
	for got := range seq {
		if got != (Token{Text, "ab"}) {
			t.Errorf("first token=%v; want Text ab", got)
		}
		break
	}
	if calls != 3 || !closed {
		t.Fatalf("calls=%d,closed=%v; want 3,true", calls, closed)
	}
}

func TestCoalesceForwarding(t *testing.T) {
	base := mustNewLexer(t, &Config{Name: "coalesce-forwarding"}, Rules{"root": {{`.`, Text, nil}}})
	lexer := Coalesce(base)
	if lexer.Config() != base.Config() {
		t.Error("config identity not preserved")
	}
	var analysed string
	lexer.SetAnalyser(func(text string) float32 {
		analysed = text
		return 0.625
	})
	if score := lexer.AnalyseText("forward this text"); score != 0.625 || analysed != "forward this text" {
		t.Errorf("AnalyseText returned %v after analysing %q", score, analysed)
	}
}
