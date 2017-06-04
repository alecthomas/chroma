package lexers

import (
	"sort"

	"github.com/danwakefield/fnmatch"

	"github.com/alecthomas/chroma"
)

type prioritisedLexers []chroma.Lexer

func (p prioritisedLexers) Len() int           { return len(p) }
func (p prioritisedLexers) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p prioritisedLexers) Less(i, j int) bool { return p[i].Config().Priority < p[j].Config().Priority }

// Registry is the global Lexer registry.
var Registry = registry{byName: map[string]chroma.Lexer{}}

type registry struct {
	Lexers []chroma.Lexer
	byName map[string]chroma.Lexer
}

// Names of all lexers, optionally including aliases.
func (r *registry) Names(withAliases bool) []string {
	out := []string{}
	for _, lexer := range r.Lexers {
		config := lexer.Config()
		out = append(out, config.Name)
		if withAliases {
			out = append(out, config.Aliases...)
		}
	}
	return out
}

// Get a Lexer by name.
func (r *registry) Get(name string) chroma.Lexer {
	lexer, ok := r.byName[name]
	if ok {
		return lexer
	}
	return Fallback
}

// Match returns all lexers matching filename.
func (r *registry) Match(filename string) []chroma.Lexer {
	lexers := prioritisedLexers{}
	for _, lexer := range r.Lexers {
		config := lexer.Config()
		for _, glob := range config.Filenames {
			if fnmatch.Match(glob, filename, 0) {
				lexers = append(lexers, lexer)
				break
			}
		}
	}
	sort.Sort(lexers)
	return lexers
}

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer, err error) chroma.Lexer {
	if err != nil {
		panic(err)
	}
	config := lexer.Config()
	Registry.byName[config.Name] = lexer
	for _, alias := range config.Aliases {
		Registry.byName[alias] = lexer
	}
	Registry.Lexers = append(Registry.Lexers, lexer)
	return lexer
}
