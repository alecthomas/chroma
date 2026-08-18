package chroma

import (
	"path/filepath"
	"slices"
	"strings"
)

var (
	ignoredSuffixes = [...]string{
		// Editor backups
		"~", ".bak", ".old", ".orig",
		// Debian and derivatives apt/dpkg/ucf backups
		".dpkg-dist", ".dpkg-old", ".ucf-dist", ".ucf-new", ".ucf-old",
		// Red Hat and derivatives rpm backups
		".rpmnew", ".rpmorig", ".rpmsave",
		// Build system input/template files
		".in",
	}
)

// LexerRegistry is a registry of Lexers.
//
// Registration is not synchronised: register all lexers (typically at init
// time) before using the registry concurrently.
type LexerRegistry struct {
	Lexers  Lexers
	byName  map[string]Lexer
	byAlias map[string]Lexer
}

// NewLexerRegistry creates a new LexerRegistry of Lexers.
func NewLexerRegistry() *LexerRegistry {
	return &LexerRegistry{
		byName:  map[string]Lexer{},
		byAlias: map[string]Lexer{},
	}
}

// Names of all lexers, optionally including aliases.
func (l *LexerRegistry) Names(withAliases bool) []string {
	out := []string{}
	for _, lexer := range l.Lexers {
		config := lexer.Config()
		out = append(out, config.Name)
		if withAliases {
			out = append(out, config.Aliases...)
		}
	}
	slices.Sort(out)
	return out
}

// Aliases of all the lexers, and skip those lexers who do not have any aliases,
// or show their name instead
func (l *LexerRegistry) Aliases(skipWithoutAliases bool) []string {
	out := []string{}
	for _, lexer := range l.Lexers {
		config := lexer.Config()
		if len(config.Aliases) == 0 {
			if skipWithoutAliases {
				continue
			}
			out = append(out, config.Name)
		}
		out = append(out, config.Aliases...)
	}
	slices.Sort(out)
	return out
}

// Get a Lexer by name, alias or file extension.
func (l *LexerRegistry) Get(name string) Lexer {
	if lexer := l.byName[name]; lexer != nil {
		return lexer
	}
	if lexer := l.byAlias[name]; lexer != nil {
		return lexer
	}
	if lexer := l.byName[strings.ToLower(name)]; lexer != nil {
		return lexer
	}
	if lexer := l.byAlias[strings.ToLower(name)]; lexer != nil {
		return lexer
	}

	candidates := PrioritisedLexers{}
	// Try file extension.
	if lexer := l.Match("filename." + name); lexer != nil {
		candidates = append(candidates, lexer)
	}
	// Try exact filename.
	if lexer := l.Match(name); lexer != nil {
		candidates = append(candidates, lexer)
	}
	if len(candidates) == 0 {
		return nil
	}
	return slices.MinFunc(candidates, compareLexersByPriority)
}

// MatchMimeType attempts to find a lexer for the given MIME type.
func (l *LexerRegistry) MatchMimeType(mimeType string) Lexer {
	matched := PrioritisedLexers{}
	for _, l := range l.Lexers {
		for _, lmt := range l.Config().MimeTypes {
			if mimeType == lmt {
				matched = append(matched, l)
			}
		}
	}
	if len(matched) != 0 {
		return slices.MinFunc(matched, compareLexersByPriority)
	}
	return nil
}

// Match returns the first lexer matching filename.
//
// Note that this iterates over all file patterns in all lexers, so is not fast.
func (l *LexerRegistry) Match(filename string) Lexer {
	filename = filepath.Base(filename)
	// First, try primary filename matches.
	matched := PrioritisedLexers{}
	for _, lexer := range l.Lexers {
		if matchGlobs(lexer.Config().Filenames, filename) {
			matched = append(matched, lexer)
		}
	}
	if len(matched) > 0 {
		return slices.MinFunc(matched, compareLexersByPriority)
	}
	// Next, try filename aliases.
	matched = nil
	for _, lexer := range l.Lexers {
		if matchGlobs(lexer.Config().AliasFilenames, filename) {
			matched = append(matched, lexer)
		}
	}
	if len(matched) > 0 {
		return slices.MinFunc(matched, compareLexersByPriority)
	}
	return nil
}

// matchGlobs reports whether filename matches any of globs, either directly
// or with one of the ignoredSuffixes appended.
func matchGlobs(globs []string, filename string) bool {
	for _, glob := range globs {
		ok, err := filepath.Match(glob, filename)
		if err != nil { // nolint
			panic(err)
		}
		if ok {
			return true
		}
		for _, suf := range &ignoredSuffixes {
			ok, err := filepath.Match(glob+suf, filename)
			if err != nil {
				panic(err)
			}
			if ok {
				return true
			}
		}
	}
	return false
}

// Analyse text content and return the "best" lexer..
func (l *LexerRegistry) Analyse(text string) Lexer {
	var picked Lexer
	highest := float32(0.0)
	for _, lexer := range l.Lexers {
		weight := lexer.AnalyseText(text)
		if weight > highest {
			picked = lexer
			highest = weight
		}
	}
	return picked
}

// Register a Lexer with the LexerRegistry. If the lexer is already registered
// it will be replaced.
//
// Not safe to call concurrently with any other method on the registry.
func (l *LexerRegistry) Register(lexer Lexer) Lexer {
	lexer.SetRegistry(l)
	config := lexer.Config()

	// Consult byName before updating it so that registering a fresh lexer,
	// the common case, does not scan the whole slice. byName is also keyed
	// on lowercased names, so a hit only means a probable duplicate.
	if l.byName[config.Name] == nil {
		l.Lexers = append(l.Lexers, lexer)
	} else {
		replaced := false
		for i, val := range l.Lexers {
			if val != nil && val.Config().Name == config.Name {
				l.Lexers[i] = lexer
				replaced = true
				break
			}
		}
		if !replaced {
			l.Lexers = append(l.Lexers, lexer)
		}
	}

	l.byName[config.Name] = lexer
	l.byName[strings.ToLower(config.Name)] = lexer

	for _, alias := range config.Aliases {
		l.byAlias[alias] = lexer
		l.byAlias[strings.ToLower(alias)] = lexer
	}

	return lexer
}
