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
type Option func(h *HTMLFormatter)

// Standalone configures the HTML formatter for generating a standalone HTML document.
func Standalone() Option { return func(h *HTMLFormatter) { h.standalone = true } }

// ClassPrefix sets the CSS class prefix.
func ClassPrefix(prefix string) Option { return func(h *HTMLFormatter) { h.prefix = prefix } }

// WithClasses emits HTML using CSS classes, rather than inline styles.
func WithClasses() Option { return func(h *HTMLFormatter) { h.classes = true } }

// New HTML formatter.
func New(options ...Option) *HTMLFormatter {
	h := &HTMLFormatter{}
	for _, option := range options {
		option(h)
	}
	return h
}

type HTMLFormatter struct {
	standalone bool
	prefix     string
	classes    bool
}

func (h *HTMLFormatter) Format(w io.Writer, style *chroma.Style) (func(*chroma.Token), error) {
	if h.classes {
		return h.formatWithClasses(w, style)
	}
	return h.formatWithoutClasses(w, style)
}

func (h *HTMLFormatter) formatWithoutClasses(w io.Writer, style *chroma.Style) (func(*chroma.Token), error) {
	classes := h.typeStyles(style)
	bg := compressStyle(classes[chroma.Background])
	if h.standalone {
		fmt.Fprint(w, "<html>\n")
		fmt.Fprintf(w, "<body style=\"%s\">\n", bg)
	}
	fmt.Fprintf(w, "<pre style=\"%s\">\n", bg)
	for t, style := range classes {
		classes[t] = compressStyle(style)
	}
	return func(token *chroma.Token) {
		if token.Type == chroma.EOF {
			fmt.Fprint(w, "</pre>\n")
			if h.standalone {
				fmt.Fprint(w, "</body>\n")
				fmt.Fprint(w, "</html>\n")
			}
			return
		}

		html := html.EscapeString(token.String())
		style := classes[token.Type]
		if style == "" {
			style = classes[token.Type.SubCategory()]
			if style == "" {
				style = classes[token.Type.Category()]
			}
		}
		if style == "" {
			fmt.Fprint(w, html)
		} else {
			fmt.Fprintf(w, "<span style=\"%s\">%s</span>", style, html)
		}
	}, nil
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

func (h *HTMLFormatter) formatWithClasses(w io.Writer, style *chroma.Style) (func(*chroma.Token), error) {
	classes := h.typeStyles(style)
	if h.standalone {
		fmt.Fprint(w, "<html>\n")
		fmt.Fprint(w, "<style type=\"text/css\">\n")
		h.WriteCSS(w, style)
		fmt.Fprintf(w, "body { %s; }\n", classes[chroma.Background])
		fmt.Fprint(w, "</style>\n")
		fmt.Fprint(w, "<body>\n")
	}
	fmt.Fprint(w, "<pre class=\"chroma\">\n")
	return func(token *chroma.Token) {
		if token.Type == chroma.EOF {
			fmt.Fprint(w, "</pre>\n")
			if h.standalone {
				fmt.Fprint(w, "</body>\n")
				fmt.Fprint(w, "</html>\n")
			}
			return
		}

		tt := token.Type
		class := classes[tt]
		if class == "" {
			tt = tt.SubCategory()
			class = classes[tt]
			if class == "" {
				tt = tt.Category()
				class = classes[tt]
			}
		}
		if class == "" {
			fmt.Fprint(w, token)
		} else {
			html := html.EscapeString(token.String())
			fmt.Fprintf(w, "<span class=\"%ss%x\">%s</span>", h.prefix, int(tt), html)
		}
	}, nil
}

// WriteCSS writes CSS style definitions (without any surrounding HTML).
func (h *HTMLFormatter) WriteCSS(w io.Writer, style *chroma.Style) error {
	classes := h.typeStyles(style)
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
		if _, err := fmt.Fprintf(w, "/* %s */ .chroma .%ss%x { %s }\n", tt, h.prefix, int(tt), styles); err != nil {
			return err
		}
	}
	return nil
}

func (h *HTMLFormatter) typeStyles(style *chroma.Style) map[chroma.TokenType]string {
	bg := style.Get(chroma.Background)
	classes := map[chroma.TokenType]string{}
	for t := range style.Entries {
		e := style.Entries[t]
		if t != chroma.Background {
			e = e.Sub(bg)
		}
		styles := h.class(e)
		classes[t] = strings.Join(styles, "; ")
	}
	return classes
}

func (h *HTMLFormatter) class(e *chroma.StyleEntry) []string {
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
	return styles
}
