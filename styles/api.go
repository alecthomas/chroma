package styles

import (
	"embed"
	"io/fs"
	"sort"

	"github.com/alecthomas/chroma/v2"
)

//go:embed *.xml
var embedded embed.FS

// Registry of Styles.
var Registry = func() map[string]*chroma.Style {
	registry := map[string]*chroma.Style{}
	// Register all embedded styles.
	files, err := fs.ReadDir(embedded, ".")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		style, err := chroma.NewEmbeddedXMLStyle(&embedded, file.Name())
		if err != nil {
			panic(err)
		}
		registry[style.Name] = style
	}
	return registry
}()

// Fallback style. Reassign to change the default fallback style.
var Fallback = Registry["swapoff"]

// Register a chroma.Style.
func Register(style *chroma.Style) *chroma.Style {
	Registry[style.Name] = style
	return style
}

// Names of all available styles.
func Names() []string {
	out := []string{}
	for name := range Registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Get named style, or Fallback.
func Get(name string) *chroma.Style {
	if style, ok := Registry[name]; ok {
		return style
	}
	return Fallback
}
