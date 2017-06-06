package chroma

import "strings"

// InheritStyle from entry with this key.
const InheritStyle TokenType = -1

// A StyleEntry in the Style map.
type StyleEntry struct {
	// Hex colours.
	Colour     Colour
	Background Colour
	Border     Colour

	Bold      bool
	Italic    bool
	Underline bool
}

func (e *StyleEntry) String() string {
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

func (e *StyleEntry) IsZero() bool {
	return e.Colour == 0 && e.Background == 0 && e.Border == 0 && !e.Bold && !e.Italic && !e.Underline
}

// StyleEntries mapping TokenType to colour definition.
type StyleEntries map[TokenType]string

// NewStyle creates a new style definition.
func NewStyle(name string, entries StyleEntries) *Style {
	s := &Style{
		Name: name,
		Entries: map[TokenType]*StyleEntry{
			InheritStyle: &StyleEntry{},
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
	Entries map[TokenType]*StyleEntry
}

// Get a style entry. Will try sub-category or category if an exact match is not found, and
// finally return the entry mapped to `InheritStyle`.
func (s *Style) Get(ttype TokenType) *StyleEntry {
	out := s.Entries[ttype]
	if out == nil {
		out = s.Entries[ttype.SubCategory()]
		if out == nil {
			out = s.Entries[ttype.Category()]
			if out == nil {
				out = s.Entries[InheritStyle]
			}
		}
	}
	return out
}

// Add an StyleEntry to the Style map.
//
// See http://pygments.org/docs/styles/#style-rules for details.
func (s *Style) Add(ttype TokenType, entry string) *Style { // nolint: gocyclo
	out := &StyleEntry{}
	dupl := s.Entries[ttype.SubCategory()]
	if dupl == nil {
		dupl = s.Entries[ttype.Category()]
		if dupl == nil {
			dupl = s.Entries[InheritStyle]
		}
	}
	parent := &StyleEntry{}
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
			parent = &StyleEntry{}
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
