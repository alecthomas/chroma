package chroma

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestLexerRegistryAllIsSortedCopy(t *testing.T) {
	registry := NewLexerRegistry()
	for _, name := range []string{"beta", "Gamma", "alpha"} {
		registry.Register(mustNewLexer(t, &Config{Name: name}, Rules{"root": {}}))
	}
	names := func(lexers []Lexer) []string {
		out := []string{}
		for _, lexer := range lexers {
			out = append(out, lexer.Config().Name)
		}
		return out
	}
	all := registry.All()
	assert.Equal(t, []string{"alpha", "beta", "Gamma"}, names(all))
	all[0] = nil
	assert.Equal(t, []string{"alpha", "beta", "Gamma"}, names(registry.All()))
}

func TestLexerRegistryReplaceAliases(t *testing.T) {
	registry := NewLexerRegistry()
	old := mustNewLexer(t, &Config{
		Name:    "Example",
		Aliases: []string{"OldAlias", "retained", "SharedAlias"},
	}, Rules{"root": {}})
	other := mustNewLexer(t, &Config{
		Name:    "Other",
		Aliases: []string{"SharedAlias"},
	}, Rules{"root": {}})
	replacement := mustNewLexer(t, &Config{
		Name:    "Example",
		Aliases: []string{"NewAlias", "retained"},
	}, Rules{"root": {}})
	registry.Register(old)
	registry.Register(other)
	registry.Register(replacement)

	for _, alias := range []string{"OldAlias", "oldalias"} {
		if got := registry.Get(alias); got != nil {
			t.Errorf("Get(%q) returned the removed lexer %q", alias, got.Config().Name)
		}
	}
	for _, name := range []string{"Example", "example", "NewAlias", "newalias", "retained"} {
		if got := registry.Get(name); got != replacement {
			t.Errorf("Get(%q) did not return the replacement lexer", name)
		}
	}
	for _, alias := range []string{"SharedAlias", "sharedalias"} {
		if got := registry.Get(alias); got != other {
			t.Errorf("Get(%q) did not preserve the other lexer's alias", alias)
		}
	}
}
