package lexers_test

import (
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
)

// configLexer stands in for a registered lexer so that registering it in a
// test registry does not repoint the real lexer's registry.
type configLexer struct {
	chroma.Lexer
	config *chroma.Config
}

func (c *configLexer) Config() *chroma.Config                         { return c.config }
func (c *configLexer) SetRegistry(*chroma.LexerRegistry) chroma.Lexer { return c }

var backupSuffixes = []string{"~", ".bak", ".old", ".orig", ".dpkg-dist", ".dpkg-old", ".ucf-dist", ".ucf-new", ".ucf-old", ".rpmnew", ".rpmorig", ".rpmsave", ".in"}

// linearMatch is the original unindexed implementation of LexerRegistry.Match.
func linearMatch(registered []chroma.Lexer, filename string) chroma.Lexer {
	filename = filepath.Base(filename)
	matches := func(globs []string) bool {
		for _, glob := range globs {
			for _, suffix := range append([]string{""}, backupSuffixes...) {
				if ok, _ := filepath.Match(glob+suffix, filename); ok {
					return true
				}
			}
		}
		return false
	}
	for _, aliases := range []bool{false, true} {
		var best chroma.Lexer
		for _, lexer := range registered {
			globs := lexer.Config().Filenames
			if aliases {
				globs = lexer.Config().AliasFilenames
			}
			if matches(globs) && (best == nil || priority(lexer) > priority(best)) {
				best = lexer
			}
		}
		if best != nil {
			return best
		}
	}
	return nil
}

func priority(lexer chroma.Lexer) float32 {
	if p := lexer.Config().Priority; p != 0 {
		return p
	}
	return 1
}

var globMeta = regexp.MustCompile(`\[([^]])[^]]*\]|\*|\?`)

// sampleFilenames returns names that match, nearly match, or extend globs.
func sampleFilenames(globs []string) []string {
	out := []string{"", "x", ".", "noextension", "a.zzz"}
	for i, glob := range globs {
		name := globMeta.ReplaceAllStringFunc(glob, func(meta string) string {
			if strings.HasPrefix(meta, "[") {
				return meta[1:2]
			}
			return "a"
		})
		suffix := backupSuffixes[i%len(backupSuffixes)]
		out = append(out, name, strings.ToUpper(name), "dir/"+name, name+"x", "x"+name, name+suffix)
	}
	slices.Sort(out)
	return slices.Compact(out)
}

func TestMatchAgreesWithLinearMatch(t *testing.T) {
	registry := chroma.NewLexerRegistry()
	registered := []chroma.Lexer{}
	globs := []string{}
	for _, lexer := range lexers.All() {
		stub := &configLexer{config: lexer.Config()}
		registry.Register(stub)
		registered = append(registered, stub)
		globs = append(globs, lexer.Config().Filenames...)
		globs = append(globs, lexer.Config().AliasFilenames...)
	}
	names := sampleFilenames(globs)
	// The linear reference takes milliseconds per name.
	want := make([]chroma.Lexer, len(names))
	var wg sync.WaitGroup
	for worker := range runtime.GOMAXPROCS(0) {
		wg.Go(func() {
			for i := worker; i < len(names); i += runtime.GOMAXPROCS(0) {
				want[i] = linearMatch(registered, names[i])
			}
		})
	}
	wg.Wait()
	for i, name := range names {
		assert.Equal(t, want[i], registry.Match(name), "Match(%q)", name)
	}
}

func BenchmarkMatch(b *testing.B) {
	for _, name := range []string{"main.go", "dir/Makefile", "Dockerfile.dev", "main.go.bak", "nosuchfile.zzz"} {
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				lexers.Match(name)
			}
		})
	}
}
