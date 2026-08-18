package styles

import (
	"embed"
	"io/fs"
	"slices"
	"strings"
	"sync"

	"github.com/alecthomas/chroma/v3"
)

//go:embed *.xml
var embedded embed.FS

var (
	registryMu sync.Mutex
	// registry holds user-registered styles plus embedded styles that have
	// already been parsed. Embedded styles are parsed lazily by Lookup so
	// that importing this package does not pay for parsing every style;
	// see TestEmbeddedStyleNamesMatchFilenames for the invariant that makes
	// the lazy filename-based lookup possible.
	registry = map[string]*chroma.Style{}
)

// embeddedNames returns the lowercased names of the embedded styles, which
// by convention are also their file names.
var embeddedNames = sync.OnceValue(func() []string {
	files, err := fs.ReadDir(embedded, ".")
	if err != nil {
		panic(err)
	}
	names := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".xml") {
			continue
		}
		names = append(names, strings.TrimSuffix(file.Name(), ".xml"))
	}
	return names
})

// Fallback style. Reassign to change the default fallback style.
var Fallback = func() *chroma.Style {
	fallback, ok := Lookup("swapoff")
	if !ok {
		panic(`chroma/styles: default fallback style "swapoff" is missing`)
	}
	return fallback
}()

// Register a chroma.Style.
func Register(style *chroma.Style) *chroma.Style {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry[strings.ToLower(style.Name)] = style
	return style
}

// Names of all available styles.
func Names() []string {
	registryMu.Lock()
	defer registryMu.Unlock()
	names := slices.Clone(embeddedNames())
	for name := range registry {
		if !slices.Contains(names, name) {
			names = append(names, name)
		}
	}
	slices.Sort(names)
	return names
}

// Lookup a named style, returning false if not found. Embedded styles are
// parsed on first lookup.
func Lookup(name string) (*chroma.Style, bool) {
	name = strings.ToLower(name)
	registryMu.Lock()
	defer registryMu.Unlock()
	if style, ok := registry[name]; ok {
		return style, true
	}
	f, err := embedded.Open(name + ".xml")
	if err != nil {
		return nil, false
	}
	defer f.Close()
	style, err := chroma.NewXMLStyle(f)
	if err != nil {
		panic(err)
	}
	registry[name] = style
	return style, true
}

// Get named style, or Fallback.
func Get(name string) *chroma.Style {
	if style, ok := Lookup(name); ok {
		return style
	}
	return Fallback
}

// GetForMode returns the named style if it already matches mode, otherwise its
// registered counterpart if one exists and matches mode. If neither matches,
// the originally-requested style is returned (or Fallback if the name is
// unknown), so callers always get something usable.
func GetForMode(name string, mode chroma.Mode) *chroma.Style {
	style := Get(name)
	if style.Mode() == mode {
		return style
	}
	if style.Counterpart == "" {
		return style
	}
	counterpart, ok := Lookup(style.Counterpart)
	if !ok || counterpart.Mode() != mode {
		return style
	}
	return counterpart
}

// RegisterPair links two styles as light/dark counterparts of each other.
//
// Both styles are also registered if they are not already present.
func RegisterPair(a, b *chroma.Style) {
	Register(a)
	Register(b)
	a.Counterpart = strings.ToLower(b.Name)
	b.Counterpart = strings.ToLower(a.Name)
}
