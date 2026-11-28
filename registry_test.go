package chroma

import (
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestLexerRegistryMatch(t *testing.T) {
	registry := NewLexerRegistry()
	register := func(config *Config) {
		registry.Register(mustNewLexer(t, config, Rules{"root": {}}))
	}
	register(&Config{Name: "first", Filenames: []string{"*.x", "Exact"}})
	register(&Config{Name: "second", Filenames: []string{"*.x", "*.y"}})
	register(&Config{Name: "preferred", Filenames: []string{"*.y"}, Priority: 2})
	register(&Config{Name: "pattern", Filenames: []string{"*.[ab]z", "pre*"}})
	register(&Config{Name: "alias", Filenames: []string{"*.p"}, AliasFilenames: []string{"*.x", "*.q"}})
	register(&Config{Name: "dotted", Filenames: []string{"*.d.ts"}})
	tests := []struct {
		filename, want string
	}{
		{"a.x", "first"},
		{"dir/a.x", "first"},
		{"a.y", "preferred"},
		{"Exact", "first"},
		{"exact", ""},
		{"a.bz", "pattern"},
		{"prefix.txt", "pattern"},
		{"a.q", "alias"},
		{"a.p", "alias"},
		{"a.d.ts", "dotted"},
		{".x", "first"},
		{"a.x.bak", "first"},
		{"a.x~", "first"},
		{"Exact.orig", "first"},
		{"a.bz.in", "pattern"},
		{"a.x.y", "preferred"},
		{"a.xx", ""},
		{"a", ""},
	}
	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			got := ""
			if lexer := registry.Match(test.filename); lexer != nil {
				got = lexer.Config().Name
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestLexerRegistryMatchAfterReplace(t *testing.T) {
	registry := NewLexerRegistry()
	registry.Register(mustNewLexer(t, &Config{Name: "a", Filenames: []string{"*.old"}}, Rules{"root": {}}))
	registry.Register(mustNewLexer(t, &Config{Name: "b", Filenames: []string{"*.new"}}, Rules{"root": {}}))
	replacement := mustNewLexer(t, &Config{Name: "a", Filenames: []string{"*.new"}}, Rules{"root": {}})
	registry.Register(replacement)
	assert.Zero(t, registry.Match("f.old"))
	assert.Equal(t, Lexer(replacement), registry.Match("f.new"), "replacement keeps the original registration position")
}

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
