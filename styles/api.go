package styles

import (
	"sort"
	"sync"

	"github.com/alecthomas/chroma/v2"
)

// Registry of Styles.
var Registry = sync.Map{}

// Fallback style. Reassign to change the default fallback style.
var Fallback = SwapOff

// Register a chroma.Style.
func Register(style *chroma.Style) *chroma.Style {
	Registry.Store(style.Name, style)
	return style
}

// Names of all available styles.
func Names() []string {
	out := []string{}
	Registry.Range(func(key, value interface{}) bool {
		k, ok := key.(string)
		if ok {
			out = append(out, k)
		}
		return true
	})
	sort.Strings(out)
	return out
}

// Get named style, or Fallback.
func Get(name string) *chroma.Style {
	if style, ok := Registry.Load(name); ok {
		if s, ok := style.(*chroma.Style); ok {
			return s
		}
	}
	return Fallback
}
