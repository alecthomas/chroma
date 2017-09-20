package chroma

import (
	"sort"
	"strings"
)

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

// Clone this StyleEntry.
func (s *StyleEntry) Clone() *StyleEntry {
	clone := &StyleEntry{}
	*clone = *s
	return clone
}

func (s *StyleEntry) String() string {
	out := []string{}
	if s.Bold {
		out = append(out, "bold")
	}
	if s.Italic {
		out = append(out, "italic")
	}
	if s.Underline {
		out = append(out, "underline")
	}
	if s.Colour.IsSet() {
		out = append(out, s.Colour.String())
	}
	if s.Background.IsSet() {
		out = append(out, "bg:"+s.Background.String())
	}
	if s.Border.IsSet() {
		out = append(out, "border:"+s.Border.String())
	}
	return strings.Join(out, " ")
}

func (s *StyleEntry) IsZero() bool {
	return s.Colour == 0 && s.Background == 0 && s.Border == 0 && !s.Bold && !s.Italic && !s.Underline
}

func (s *StyleEntry) Sub(e *StyleEntry) *StyleEntry {
	out := &StyleEntry{}
	if e.Colour != s.Colour {
		out.Colour = s.Colour
	}
	if e.Background != s.Background {
		out.Background = s.Background
	}
	if e.Bold != s.Bold {
		out.Bold = s.Bold
	}
	if e.Italic != s.Italic {
		out.Italic = s.Italic
	}
	if e.Underline != s.Underline {
		out.Underline = s.Underline
	}
	if e.Border != s.Border {
		out.Border = s.Border
	}
	return out
}

// StyleEntries mapping TokenType to colour definition.
type StyleEntries map[TokenType]string

// NewStyle creates a new style definition.
func NewStyle(name string, entries StyleEntries) *Style {
	s := &Style{
		Name:    name,
		Entries: map[TokenType]*StyleEntry{},
	}
	s.Add(Background, "")
	s.AddAll(entries)
	return s
}

// A Style definition.
//
// See http://pygments.org/docs/styles/ for details. Semantics are intended to be identical.
type Style struct {
	Name    string
	Entries map[TokenType]*StyleEntry
}

// Clone this style. The clone can then be safely modified.
func (s *Style) Clone() *Style {
	clone := &Style{
		Name:    s.Name,
		Entries: map[TokenType]*StyleEntry{},
	}
	for tt, e := range s.Entries {
		clone.Entries[tt] = e.Clone()
	}
	return clone
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
				out = s.Entries[Background]
			}
		}
	}
	return out
}

func (s *Style) AddAll(entries StyleEntries) *Style {
	tis := []int{}
	for tt := range entries {
		tis = append(tis, int(tt))
	}
	sort.Ints(tis)
	for _, ti := range tis {
		tt := TokenType(ti)
		entry := entries[tt]
		s.Add(tt, entry)
	}
	return s
}

// Add a StyleEntry to the Style map.
//
// See http://pygments.org/docs/styles/#style-rules for details.
func (s *Style) Add(ttype TokenType, entry string) *Style { // nolint: gocyclo
	dupl := s.Entries[ttype.SubCategory()]
	if dupl == nil {
		dupl = s.Entries[ttype.Category()]
		if dupl == nil {
			dupl = s.Entries[Background]
			if dupl == nil {
				dupl = &StyleEntry{}
			}
		}
	}
	parent := &StyleEntry{}
	// Duplicate ancestor node.
	*parent = *dupl
	s.Entries[ttype] = ParseStyleEntry(parent, entry)
	return s
}

// ParseStyleEntry parses a Pygments style entry.
func ParseStyleEntry(parent *StyleEntry, entry string) *StyleEntry { // nolint: gocyclo
	out := &StyleEntry{}
	parts := strings.Fields(entry)
	// Check if parent style should be inherited...
	if parent != nil {
		inherit := true
		for _, part := range parts {
			if part == "noinherit" {
				inherit = false
				break
			}
		}
		if inherit {
			*out = *parent
		}
	}
	for _, part := range parts {
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
		case part == "bg:":
			out.Background = 0
		case strings.HasPrefix(part, "bg:#"):
			out.Background = ParseColour(part[3:])
		case strings.HasPrefix(part, "border:#"):
			out.Border = ParseColour(part[7:])
		case strings.HasPrefix(part, "#"):
			out.Colour = ParseColour(part)
		default:
			// Here lies an error, but we ignore it in the interests of convenience.
		}
	}
	return out
}
