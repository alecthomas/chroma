package formatters

import (
	"io"
	"iter"
	"maps"
	"slices"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/formatters/svg"
)

var (
	// NoOp formatter.
	NoOp = Register("noop", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style, iterator iter.Seq[chroma.Token]) error {
		for t := range iterator {
			if _, err := io.WriteString(w, t.Value); err != nil {
				return err
			}
		}
		return nil
	}))
	// Default HTML formatter outputs self-contained HTML.
	htmlFull = Register("html", html.New(html.Standalone(true), html.WithClasses(true))) // nolint
	SVG      = Register("svg", svg.New(svg.EmbedFont("Liberation Mono", svg.FontLiberationMono, svg.WOFF)))
)

// Fallback formatter.
var Fallback = NoOp

var registry = map[string]chroma.Formatter{}

// Names of registered formatters.
func Names() []string {
	return slices.Sorted(maps.Keys(registry))
}

// Get formatter by name.
//
// If the given formatter is not found, the Fallback formatter will be returned.
func Get(name string) chroma.Formatter {
	if f, ok := registry[name]; ok {
		return f
	}
	return Fallback
}

// Register a named formatter.
//
// Not safe to call concurrently with other functions in this package;
// register all formatters at init time.
func Register(name string, formatter chroma.Formatter) chroma.Formatter {
	registry[name] = formatter
	return formatter
}
