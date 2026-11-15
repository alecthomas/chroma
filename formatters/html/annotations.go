package html

import (
	"fmt"
	"html"
	"io"
	"slices"
	"sort"
	"unicode/utf8"

	"github.com/alecthomas/chroma/v3"
)

// Annotation attaches metadata to a half-open byte range in the token text.
type Annotation struct {
	Start, End       int
	ID, Class, Title string
}

type annotationEvent struct {
	offset, annotation int
	start              bool
}

type annotationRenderer struct {
	annotations  []Annotation
	events       []annotationEvent
	next, offset int
	active       []int
}

func newAnnotationRenderer(annotations []Annotation, tokens []chroma.Token) (*annotationRenderer, error) {
	boundaries := map[int]bool{0: true}
	length := 0
	for _, token := range tokens {
		if !utf8.ValidString(token.Value) {
			return nil, fmt.Errorf("annotated token text is not valid UTF-8")
		}
		for position := range token.Value {
			boundaries[length+position] = true
		}
		length += len(token.Value)
		boundaries[length] = true
	}
	renderer := &annotationRenderer{annotations: annotations}
	ids := map[string]bool{}
	for index, annotation := range annotations {
		if annotation.ID == "" || ids[annotation.ID] {
			return nil, fmt.Errorf("annotation IDs must be nonempty and unique")
		}
		ids[annotation.ID] = true
		if annotation.Start < 0 || annotation.End <= annotation.Start || annotation.End > length {
			return nil, fmt.Errorf("annotation %q has invalid range [%d,%d)", annotation.ID, annotation.Start, annotation.End)
		}
		if !boundaries[annotation.Start] || !boundaries[annotation.End] {
			return nil, fmt.Errorf("annotation %q splits a UTF-8 character", annotation.ID)
		}
		renderer.events = append(renderer.events,
			annotationEvent{offset: annotation.Start, annotation: index, start: true},
			annotationEvent{offset: annotation.End, annotation: index})
	}
	sort.Slice(renderer.events, func(i, j int) bool { return renderer.events[i].offset < renderer.events[j].offset })
	return renderer, nil
}

func (r *annotationRenderer) advance() {
	for r.next < len(r.events) && r.events[r.next].offset <= r.offset {
		event := r.events[r.next]
		position, found := slices.BinarySearch(r.active, event.annotation)
		if event.start && !found {
			r.active = slices.Insert(r.active, position, event.annotation)
		} else if !event.start && found {
			r.active = slices.Delete(r.active, position, position+1)
		}
		r.next++
	}
}

func (r *annotationRenderer) writeToken(f *Formatter, w io.Writer, css map[chroma.TokenType]string, token chroma.Token) {
	text := token.Value
	for len(text) > 0 {
		r.advance()
		length := len(text)
		if r.next < len(r.events) {
			length = min(length, r.events[r.next].offset-r.offset)
		}
		for _, index := range r.active {
			annotation := r.annotations[index]
			fmt.Fprintf(w, `<span data-annotation="%s"`, html.EscapeString(annotation.ID))
			if annotation.Class != "" {
				fmt.Fprintf(w, ` class="%s"`, html.EscapeString(annotation.Class))
			}
			if annotation.Title != "" {
				fmt.Fprintf(w, ` title="%s"`, html.EscapeString(annotation.Title))
			}
			fmt.Fprint(w, ">")
		}
		attr := f.styleAttr(css, token.Type)
		if attr != "" {
			fmt.Fprintf(w, "<span%s>", attr)
		}
		fmt.Fprint(w, html.EscapeString(text[:length]))
		if attr != "" {
			fmt.Fprint(w, "</span>")
		}
		for range r.active {
			fmt.Fprint(w, "</span>")
		}
		text = text[length:]
		r.offset += length
	}
}

type annotationWriter struct {
	writer io.Writer
	err    error
}

func (w *annotationWriter) Write(data []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	n, err := w.writer.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}
	w.err = err
	return n, err
}
