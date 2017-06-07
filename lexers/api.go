package lexers

import (
	"path/filepath"

	"github.com/danwakefield/fnmatch"

	"github.com/alecthomas/chroma"
)

// Registry of Lexers.
var Registry = struct {
	Lexers []chroma.Lexer
	byName map[string]chroma.Lexer
}{
	byName: map[string]chroma.Lexer{},
}

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string {
	out := []string{}
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		out = append(out, config.Name)
		if withAliases {
			out = append(out, config.Aliases...)
		}
	}
	return out
}

// Get a Lexer by name.
func Get(name string) chroma.Lexer {
	lexer, ok := Registry.byName[name]
	if ok {
		return lexer
	}
	return nil
}

// Match returns all lexers matching filename.
func Match(filename string) chroma.Lexer {
	filename = filepath.Base(filename)
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.Filenames {
			if fnmatch.Match(glob, filename, 0) {
				return lexer
			}
		}
	}
	return nil
}

// Analyse text content and return the "best" lexer..
func Analyse(text string) chroma.Lexer {
	var picked chroma.Lexer
	highest := float32(0.0)
	for _, lexer := range Registry.Lexers {
		if analyser, ok := lexer.(chroma.Analyser); ok {
			weight := analyser.AnalyseText(text)
			if weight > highest {
				picked = lexer
				highest = weight
			}
		}
	}
	return picked
}

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer) chroma.Lexer {
	config := lexer.Config()
	Registry.byName[config.Name] = lexer
	for _, alias := range config.Aliases {
		Registry.byName[alias] = lexer
	}
	Registry.Lexers = append(Registry.Lexers, lexer)
	return lexer
}
