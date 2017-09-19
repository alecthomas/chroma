package html

import (
	"fmt"
	"html"
	"io"
	"sort"
	"strings"

	"github.com/alecthomas/chroma"
)

// Option sets an option of the HTML formatter.
type Option func(f *Formatter)

// Standalone configures the HTML formatter for generating a standalone HTML document.
func Standalone() Option { return func(f *Formatter) { f.standalone = true } }

// ClassPrefix sets the CSS class prefix.
func ClassPrefix(prefix string) Option { return func(f *Formatter) { f.prefix = prefix } }

// WithClasses emits HTML using CSS classes, rather than inline styles.
func WithClasses() Option { return func(f *Formatter) { f.classes = true } }

// TabWidth sets the number of characters for a tab. Defaults to 8.
func TabWidth(width int) Option { return func(f *Formatter) { f.tabWidth = width } }

// WithLineNumbers formats output with line numbers.
func WithLineNumbers() Option {
	return func(f *Formatter) {
		f.lineNumbers = true
	}
}

// HighlightLines higlights the given line ranges.
//
// A range is the beginning and ending of a range as 1-based line numbers, inclusive.
func HighlightLines(style string, ranges [][2]int) Option {
	return func(f *Formatter) {
		f.highlightStyle = style
		f.highlightRanges = ranges
	}
}

// New HTML formatter.
func New(options ...Option) *Formatter {
	f := &Formatter{}
	for _, option := range options {
		option(f)
	}
	return f
}

// Formatter that generates HTML.
type Formatter struct {
	standalone      bool
	prefix          string
	classes         bool
	tabWidth        int
	lineNumbers     bool
	highlightStyle  string
	highlightRanges [][2]int
}

func (f *Formatter) Format(w io.Writer, style *chroma.Style) (func(*chroma.Token), error) {
	styles := f.typeStyles(style)
	if !f.classes {
		for t, style := range styles {
			styles[t] = compressStyle(style)
		}
	}
	if f.standalone {
		fmt.Fprint(w, "<html>\n")
		if f.classes {
			fmt.Fprint(w, "<style type=\"text/css\">\n")
			f.WriteCSS(w, style)
			fmt.Fprintf(w, "body { %s; }\n", styles[chroma.Background])
			fmt.Fprint(w, "</style>")
		}
		fmt.Fprintf(w, "<body%s>\n", f.styleAttr(styles, chroma.Background))
	}
	fmt.Fprintf(w, "<pre%s>\n", f.styleAttr(styles, chroma.Background))
	return func(token *chroma.Token) {
		if token.Type == chroma.EOF {
			fmt.Fprint(w, "</pre>\n")
			if f.standalone {
				fmt.Fprint(w, "</body>\n")
				fmt.Fprint(w, "</html>\n")
			}
			return
		}
		html := html.EscapeString(token.String())
		attr := f.styleAttr(styles, token.Type)
		if attr == "" {
			fmt.Fprint(w, html)
		} else {
			fmt.Fprintf(w, "<span%s>%s</span>", attr, html)
		}
	}, nil
}

func (f *Formatter) class(tt chroma.TokenType) string {
	if tt == chroma.Background {
		return "chroma"
	}
	if tt < 0 {
		return fmt.Sprintf("%sss%x", f.prefix, -int(tt))
	}
	return fmt.Sprintf("%ss%x", f.prefix, int(tt))
}

func (f *Formatter) styleAttr(styles map[chroma.TokenType]string, tt chroma.TokenType) string {
	if _, ok := styles[tt]; !ok {
		tt = tt.SubCategory()
		if _, ok := styles[tt]; !ok {
			tt = tt.Category()
			if _, ok := styles[tt]; !ok {
				return ""
			}
		}
	}
	if f.classes {
		return string(fmt.Sprintf(` class="%s"`, f.class(tt)))
	}
	return string(fmt.Sprintf(` style="%s"`, styles[tt]))
}

func (f *Formatter) tabWidthStyle() string {
	if f.tabWidth != 0 && f.tabWidth != 8 {
		return fmt.Sprintf("; -moz-tab-size: %[1]d; -o-tab-size: %[1]d; tab-size: %[1]d", f.tabWidth)
	}
	return ""
}

func compressStyle(s string) string {
	s = strings.Replace(s, " ", "", -1)
	parts := strings.Split(s, ";")
	out := []string{}
	for _, p := range parts {
		if strings.Contains(p, "#") {
			c := p[len(p)-6:]
			if c[0] == c[1] && c[2] == c[3] && c[4] == c[5] {
				p = p[:len(p)-6] + c[0:1] + c[2:3] + c[4:5]
			}
		}
		out = append(out, p)
	}
	return strings.Join(out, ";")
}

// WriteCSS writes CSS style definitions (without any surrounding HTML).
func (f *Formatter) WriteCSS(w io.Writer, style *chroma.Style) error {
	classes := f.typeStyles(style)
	if _, err := fmt.Fprintf(w, "/* %s */ .chroma { %s }\n", chroma.Background, classes[chroma.Background]); err != nil {
		return err
	}
	tts := []int{}
	for tt := range classes {
		tts = append(tts, int(tt))
	}
	sort.Ints(tts)
	for _, ti := range tts {
		tt := chroma.TokenType(ti)
		styles := classes[tt]
		if tt < 0 {
			continue
		}
		if _, err := fmt.Fprintf(w, "/* %s */ .chroma .%ss%x { %s }\n", tt, f.prefix, int(tt), styles); err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) typeStyles(style *chroma.Style) map[chroma.TokenType]string {
	bg := style.Get(chroma.Background)
	classes := map[chroma.TokenType]string{}
	for t := range style.Entries {
		e := style.Entries[t]
		if t != chroma.Background {
			e = e.Sub(bg)
		}
		classes[t] = f.styleEntryToCSS(e)
	}
	classes[chroma.Background] += f.tabWidthStyle()
	return classes
}

func (f *Formatter) styleEntryToCSS(e *chroma.StyleEntry) string {
	styles := []string{}
	if e.Colour.IsSet() {
		styles = append(styles, "color: "+e.Colour.String())
	}
	if e.Background.IsSet() {
		styles = append(styles, "background-color: "+e.Background.String())
	}
	if e.Bold {
		styles = append(styles, "font-weight: bold")
	}
	if e.Italic {
		styles = append(styles, "font-style: italic")
	}
	return strings.Join(styles, "; ")
}
