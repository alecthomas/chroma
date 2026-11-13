package html

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"testing"

	assert "github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/styles"
)

type failingHTMLWriter struct {
	buf                bytes.Buffer
	remaining          int
	err                error
	failed             bool
	writesAfterFailure int
}

func (w *failingHTMLWriter) Write(p []byte) (int, error) {
	if w.failed {
		w.writesAfterFailure++
		return 0, errors.New("another write after failure")
	}
	if len(p) > w.remaining {
		n, _ := w.buf.Write(p[:w.remaining])
		w.failed = true
		return n, w.err
	}
	w.remaining -= len(p)
	return w.buf.Write(p)
}

func TestHTMLWriterErrors(t *testing.T) {
	tokens := []chroma.Token{
		{Type: chroma.Keyword, Value: "first & <line>\n"},
		{Type: chroma.LiteralString, Value: "second line"},
	}
	cases := []struct {
		name    string
		options []Option
	}{
		{"default", nil},
		{"classes", []Option{WithClasses(true)}},
		{"standalone inline", []Option{Standalone(true)}},
		{"standalone classes", []Option{Standalone(true), WithClasses(true), WithModeClasses(true)}},
		{"line numbers", []Option{WithLineNumbers(true), WithLinkableLineNumbers(true, "line-")}},
		{"line table", []Option{WithLineNumbers(true), LineNumbersInTable(true), HighlightLines([][2]int{{1, 1}})}},
		{"line prompts", []Option{WithLinePrompts("$ ", [][2]int{{1, 2}})}},
		{"without pre", []Option{PreventSurroundingPre(true)}},
		{"inline code", []Option{InlineCode(true)}},
		{"custom wrapper", []Option{WithPreWrapper(preWrapper{
			start: func(bool, string) string { return "<section>" },
			end:   func(bool) string { return "</section>" },
		})}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := New(tc.options...)
			var output bytes.Buffer
			assert.NoError(t, f.Format(&output, styles.Fallback, slices.Values(tokens)))
			for _, limit := range []int{0, output.Len() / 2, output.Len() - 1} {
				t.Run(fmt.Sprint(limit), func(t *testing.T) {
					writeErr := errors.New("output unavailable")
					w := &failingHTMLWriter{remaining: limit, err: writeErr}
					err := f.Format(w, styles.Fallback, slices.Values(tokens))
					if err != writeErr {
						t.Errorf("Format returned %v; want the original write error %v", err, writeErr)
					}
					assert.Equal(t, output.String()[:limit], w.buf.String())
					assert.Equal(t, 0, w.writesAfterFailure)
				})
			}
		})
	}
}

func TestHTMLWriterErrorWithEmptyInput(t *testing.T) {
	writeErr := errors.New("output unavailable")
	w := &failingHTMLWriter{err: writeErr}
	err := New().Format(w, styles.Fallback, slices.Values([]chroma.Token{}))
	if err != writeErr {
		t.Errorf("Format returned %v; want the original write error %v", err, writeErr)
	}
	assert.Equal(t, 0, w.writesAfterFailure)
}
