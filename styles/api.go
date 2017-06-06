package styles

import (
	"strings"

	"github.com/alecthomas/chroma"
)

// Registry of Styles.
var Registry = map[string]*Style{}

// Fallback style. Reassign to change the default fallback style.
var Fallback = SwapOff

// Register a Style.
func Register(style *Style) *Style {
	Registry[style.Name] = style
	return style
}

// Get named style, or Fallback.
func Get(name string) *Style {
	if style, ok := Registry[name]; ok {
		return style
	}
	return Fallback
}

// Inherit from entry with this key.
const Inherit chroma.TokenType = -1

// An Entry in the Style map.
type Entry struct {
	// Hex colours.
	Colour     Colour
	Background Colour
	Border     Colour

	Bold      bool
	Italic    bool
	Underline bool
}

func (e *Entry) String() string {
	out := []string{}
	if e.Bold {
		out = append(out, "bold")
	}
	if e.Italic {
		out = append(out, "italic")
	}
	if e.Underline {
		out = append(out, "underline")
	}
	if e.Colour.IsSet() {
		out = append(out, e.Colour.String())
	}
	if e.Background.IsSet() {
		out = append(out, "bg:"+e.Background.String())
	}
	if e.Border.IsSet() {
		out = append(out, "border:"+e.Border.String())
	}
	return strings.Join(out, " ")
}

// Entries mapping TokenType to colour definition.
type Entries map[chroma.TokenType]string

func (e *Entry) IsZero() bool {
	return e.Colour == 0 && e.Background == 0 && e.Border == 0 && !e.Bold && !e.Italic && !e.Underline
}

// New creates a new style definition.
func New(name string, entries Entries) *Style {
	s := &Style{
		Name: name,
		Entries: map[chroma.TokenType]*Entry{
			Inherit: &Entry{},
		},
	}
	for tt, entry := range entries {
		s.Add(tt, entry)
	}
	return s
}

// A Style definition.
//
// See http://pygments.org/docs/styles/ for details. Semantics are intended to be identical.
type Style struct {
	Name    string
	Entries map[chroma.TokenType]*Entry
}

// Get a style entry. Will try sub-category or category if an exact match is not found, and
// finally return the entry mapped to `Inherit`.
func (s *Style) Get(ttype chroma.TokenType) *Entry {
	out := s.Entries[ttype]
	if out == nil {
		out = s.Entries[ttype.SubCategory()]
		if out == nil {
			out = s.Entries[ttype.Category()]
			if out == nil {
				out = s.Entries[Inherit]
			}
		}
	}
	return out
}

// Add an Entry to the Style map.
//
// See http://pygments.org/docs/styles/#style-rules for details.
func (s *Style) Add(ttype chroma.TokenType, entry string) *Style { // nolint: gocyclo
	out := &Entry{}
	dupl := s.Entries[ttype.SubCategory()]
	if dupl == nil {
		dupl = s.Entries[ttype.Category()]
		if dupl == nil {
			dupl = s.Entries[Inherit]
		}
	}
	parent := &Entry{}
	// Duplicate ancestor node.
	*parent = *dupl
	for _, part := range strings.Fields(entry) {
		switch {
		case part == "italic":
			out.Italic = true
		case part == "noitalic":
			out.Italic = false
		case part == "bold":
			out.Bold = true
		case part == "nobold":
			out.Bold = false
		case part == "underline":
			out.Underline = true
		case part == "nounderline":
			out.Underline = false
		case part == "noinherit":
			parent = &Entry{}
		case strings.HasPrefix(part, "bg:#"):
			out.Background = ParseColour(part[3:])
		case strings.HasPrefix(part, "border:#"):
			out.Border = ParseColour(part[7:])
		case strings.HasPrefix(part, "#"):
			out.Colour = ParseColour(part)
		default:
			panic("unsupported style entry " + part)
		}
	}
	if parent.Colour != 0 && out.Colour == 0 {
		out.Colour = parent.Colour
	}
	if parent.Background != 0 && out.Background == 0 {
		out.Background = parent.Background
	}
	if parent.Border != 0 && out.Border == 0 {
		out.Border = parent.Border
	}
	if parent.Bold && !out.Bold {
		out.Bold = true
	}
	if parent.Italic && !out.Italic {
		out.Italic = true
	}
	if parent.Underline && !out.Underline {
		out.Underline = true
	}
	s.Entries[ttype] = out
	return s
}
