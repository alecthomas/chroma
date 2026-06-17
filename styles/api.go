package styles

import (
	"embed"
	"io/fs"
	"sort"
	"strings"

	"github.com/alecthomas/chroma/v3"
)

//go:embed *.xml
var embedded embed.FS

var registry = func() map[string]*chroma.Style {
	r := map[string]*chroma.Style{}
	files, err := fs.ReadDir(embedded, ".")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := embedded.Open(file.Name())
		if err != nil {
			panic(err)
		}
		style, err := chroma.NewXMLStyle(f)
		if err != nil {
			panic(err)
		}
		r[strings.ToLower(style.Name)] = style
		_ = f.Close()
	}
	return r
}()

// Fallback style. Reassign to change the default fallback style.
var Fallback = registry["swapoff"]

// Register a chroma.Style.
func Register(style *chroma.Style) *chroma.Style {
	registry[strings.ToLower(style.Name)] = style
	return style
}

// Names of all available styles.
func Names() []string {
	out := []string{}
	for name := range registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Lookup a named style, returning false if not found.
func Lookup(name string) (*chroma.Style, bool) {
	style, ok := registry[strings.ToLower(name)]
	return style, ok
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
	counterpart, ok := registry[style.Counterpart]
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
