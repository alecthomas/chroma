package chroma

import "testing"

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
