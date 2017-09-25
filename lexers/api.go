package lexers

import (
	"path/filepath"
	"sort"

	"github.com/danwakefield/fnmatch"

	"github.com/alecthomas/chroma"
)

// Registry of Lexers.
var Registry = struct {
	Lexers  chroma.Lexers
	byName  map[string]chroma.Lexer
	byAlias map[string]chroma.Lexer
}{
	byName:  map[string]chroma.Lexer{},
	byAlias: map[string]chroma.Lexer{},
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
	sort.Strings(out)
	return out
}

// Get a Lexer by name.
func Get(name string) chroma.Lexer {
	if lexer := Registry.byName[name]; lexer != nil {
		return lexer
	}
	return Registry.byAlias[name]
}

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) chroma.Lexer {
	for _, l := range Registry.Lexers {
		for _, lmt := range l.Config().MimeTypes {
			if mimeType == lmt {
				return l
			}
		}
	}
	return nil
}

// Match returns the first lexer matching filename.
func Match(filename string) chroma.Lexer {
	filename = filepath.Base(filename)
	// First, try primary filename matches.
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.Filenames {
			if fnmatch.Match(glob, filename, 0) {
				return lexer
			}
		}
	}
	// Next, try filename aliases.
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.AliasFilenames {
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
		Registry.byAlias[alias] = lexer
	}
	Registry.Lexers = append(Registry.Lexers, lexer)
	return lexer
}

// Fallback lexer if no other is found.
var Fallback chroma.Lexer = chroma.MustNewLexer(&chroma.Config{
	Name:      "fallback",
	Filenames: []string{"*"},
}, chroma.Rules{
	"root": []chroma.Rule{
		{`.+`, chroma.Text, nil},
		{`\n`, chroma.Text, nil},
	},
})
