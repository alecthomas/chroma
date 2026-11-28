package chroma

import (
	"cmp"
	"path/filepath"
	"slices"
	"strings"
	"sync"
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
	lexers  []Lexer
	byName  map[string]Lexer
	byAlias map[string]Lexer
	// Indexes of each lexer's Filenames and AliasFilenames, built by the
	// first Match after a Register, so later Config changes are not seen.
	globsOnce      sync.Once
	filenames      globIndex
	aliasFilenames globIndex
}

// NewLexerRegistry creates a new LexerRegistry of Lexers.
func NewLexerRegistry() *LexerRegistry {
	return &LexerRegistry{
		byName:  map[string]Lexer{},
		byAlias: map[string]Lexer{},
	}
}

// All returns a copy of the registered lexers, sorted case-insensitively by
// name.
func (l *LexerRegistry) All() []Lexer {
	return slices.SortedFunc(slices.Values(l.lexers), func(a, b Lexer) int {
		return cmp.Compare(strings.ToLower(a.Config().Name), strings.ToLower(b.Config().Name))
	})
}

// Names of all lexers, optionally including aliases.
func (l *LexerRegistry) Names(withAliases bool) []string {
	out := []string{}
	for _, lexer := range l.lexers {
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
	for _, lexer := range l.lexers {
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

	candidates := []Lexer{}
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
	matched := []Lexer{}
	for _, l := range l.lexers {
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

// Match returns the highest priority lexer whose Filenames match the base name
// of filename, falling back to its AliasFilenames. Globs also match with one
// of the ignored backup suffixes, such as ".bak", appended.
func (l *LexerRegistry) Match(filename string) Lexer {
	l.globsOnce.Do(func() {
		l.filenames = newGlobIndex(l.lexers, func(c *Config) []string { return c.Filenames })
		l.aliasFilenames = newGlobIndex(l.lexers, func(c *Config) []string { return c.AliasFilenames })
	})
	filename = filepath.Base(filename)
	if lexer := l.filenames.best(l.lexers, filename); lexer != nil {
		return lexer
	}
	return l.aliasFilenames.best(l.lexers, filename)
}

// globIndex finds the lexers, by position in LexerRegistry.lexers, whose
// filename globs match a base name.
type globIndex struct {
	// exts holds globs of the form "*.ext", keyed by ".ext" and sorted.
	exts []globEntry
	// names holds globs without metacharacters, sorted.
	names    []globEntry
	patterns []globEntry
}

type globEntry struct {
	glob  string
	lexer int
}

type globKind int

const (
	globName globKind = iota
	globExt
	globPattern
)

// classifyGlob returns how glob is indexed, and the key it is indexed by.
func classifyGlob(glob string) (globKind, string) {
	switch {
	case !hasGlobMeta(glob):
		return globName, glob
	case strings.HasPrefix(glob, "*.") && !hasGlobMeta(glob[1:]):
		return globExt, glob[1:]
	default:
		return globPattern, glob
	}
}

func hasGlobMeta(glob string) bool {
	return strings.ContainsAny(glob, `*?[\`)
}

func newGlobIndex(lexers []Lexer, globs func(*Config) []string) globIndex {
	var counts [3]int
	for _, lexer := range lexers {
		for _, glob := range globs(lexer.Config()) {
			kind, _ := classifyGlob(glob)
			counts[kind]++
		}
	}
	entries := [3][]globEntry{
		make([]globEntry, 0, counts[globName]),
		make([]globEntry, 0, counts[globExt]),
		make([]globEntry, 0, counts[globPattern]),
	}
	for i, lexer := range lexers {
		for _, glob := range globs(lexer.Config()) {
			kind, key := classifyGlob(glob)
			entries[kind] = append(entries[kind], globEntry{key, i})
		}
	}
	byGlob := func(a, b globEntry) int { return cmp.Compare(a.glob, b.glob) }
	slices.SortFunc(entries[globName], byGlob)
	slices.SortFunc(entries[globExt], byGlob)
	return globIndex{names: entries[globName], exts: entries[globExt], patterns: entries[globPattern]}
}

// lookup returns the entries in sorted whose glob is key, without
// slices.BinarySearchFunc as that makes key escape and allocate.
func lookup(sorted []globEntry, key string) []globEntry {
	start, end := 0, len(sorted)
	for start < end {
		if mid := int(uint(start+end) >> 1); sorted[mid].glob < key {
			start = mid + 1
		} else {
			end = mid
		}
	}
	end = start
	for end < len(sorted) && sorted[end].glob == key {
		end++
	}
	return sorted[start:end]
}

// best returns the highest priority lexer matching filename, preferring the
// earliest registered on a tie.
func (g *globIndex) best(lexers []Lexer, filename string) Lexer {
	best := -1
	consider := func(i int) {
		if best < 0 {
			best = i
			return
		}
		if order := compareLexersByPriority(lexers[i], lexers[best]); order < 0 || order == 0 && i < best {
			best = i
		}
	}
	match := func(name string) {
		for _, entry := range lookup(g.names, name) {
			consider(entry.lexer)
		}
		for dot := range len(name) {
			if name[dot] == '.' {
				for _, entry := range lookup(g.exts, name[dot:]) {
					consider(entry.lexer)
				}
			}
		}
		for _, pattern := range g.patterns {
			ok, err := filepath.Match(pattern.glob, name)
			if err != nil {
				panic(err)
			}
			if ok {
				consider(pattern.lexer)
			}
		}
	}
	match(filename)
	// A glob followed by a literal suffix matches exactly the names with that
	// suffix whose remainder matches the glob.
	for _, suffix := range &ignoredSuffixes {
		if name, ok := strings.CutSuffix(filename, suffix); ok {
			match(name)
		}
	}
	if best < 0 {
		return nil
	}
	return lexers[best]
}

// Analyse text content and return the "best" lexer..
func (l *LexerRegistry) Analyse(text string) Lexer {
	var picked Lexer
	highest := float32(0.0)
	for _, lexer := range l.lexers {
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
	replaced := false
	if l.byName[config.Name] != nil {
		for i, val := range l.lexers {
			if val != nil && val.Config().Name == config.Name {
				for alias, registered := range l.byAlias {
					if registered.Config().Name == config.Name {
						delete(l.byAlias, alias)
					}
				}
				l.lexers[i] = lexer
				replaced = true
				break
			}
		}
	}
	if !replaced {
		l.lexers = append(l.lexers, lexer)
	}
	// Resetting is safe only because Register must not run concurrently with
	// Match.
	l.globsOnce = sync.Once{}

	l.byName[config.Name] = lexer
	l.byName[strings.ToLower(config.Name)] = lexer

	for _, alias := range config.Aliases {
		l.byAlias[alias] = lexer
		l.byAlias[strings.ToLower(alias)] = lexer
	}

	return lexer
}
