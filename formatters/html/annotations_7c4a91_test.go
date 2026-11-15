//go:build challenge

package html_test

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"iter"
	"reflect"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/alecthomas/chroma/v3"
	html "github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/styles"
)

type annotationCharacter struct {
	value rune
	ids   []string
}
type annotationDocument struct {
	text       string
	covered    []annotationCharacter
	attributes map[string]map[string]string
	marked     map[string]string
}

func inspectAnnotations(t *testing.T, output string) annotationDocument {
	t.Helper()
	decoder := xml.NewDecoder(strings.NewReader(output))
	decoder.Entity = xml.HTMLEntity
	doc := annotationDocument{attributes: map[string]map[string]string{}, marked: map[string]string{}}
	stack := []string{}
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("invalid nested markup: %v\n%s", err, output)
		}
		switch token := token.(type) {
		case xml.StartElement:
			attrs := map[string]string{}
			for _, attr := range token.Attr {
				attrs[attr.Name.Local] = attr.Value
			}
			id := attrs["data-annotation"]
			stack = append(stack, id)
			if id != "" {
				if token.Name.Local != "span" {
					t.Fatal("annotation is not a span")
				}
				if prior, ok := doc.attributes[id]; ok && !reflect.DeepEqual(prior, attrs) {
					t.Fatal("fragment metadata changed")
				}
				doc.attributes[id] = attrs
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			value := string(token)
			doc.text += value
			active := []string{}
			for _, id := range stack {
				if id != "" {
					active = append(active, id)
					doc.marked[id] += value
				}
			}
			if len(active) != 0 {
				for _, character := range value {
					doc.covered = append(doc.covered, annotationCharacter{character, slices.Clone(active)})
				}
			}
		}
	}
	return doc
}

type sourceAnnotation struct {
	Start, End       int
	ID, Class, Title string
}

func formatAnnotated(t *testing.T, f *html.Formatter, w io.Writer, style *chroma.Style, iterator iter.Seq[chroma.Token], annotations []sourceAnnotation) error {
	t.Helper()
	method := reflect.ValueOf(f).MethodByName("FormatAnnotated")
	if !method.IsValid() {
		t.Fatal("Formatter.FormatAnnotated is missing")
	}
	typ := method.Type()
	if typ.NumIn() != 4 || typ.NumOut() != 1 || typ.In(3).Kind() != reflect.Slice || typ.In(3).Elem().Kind() != reflect.Struct {
		t.Fatal("FormatAnnotated has an incompatible public signature")
	}
	input := reflect.MakeSlice(typ.In(3), len(annotations), len(annotations))
	for i, annotation := range annotations {
		value := input.Index(i)
		for name, number := range map[string]int{"Start": annotation.Start, "End": annotation.End} {
			field := value.FieldByName(name)
			if !field.IsValid() || field.Kind() != reflect.Int {
				t.Fatalf("Annotation.%s must be int", name)
			}
			field.SetInt(int64(number))
		}
		for name, text := range map[string]string{"ID": annotation.ID, "Class": annotation.Class, "Title": annotation.Title} {
			field := value.FieldByName(name)
			if !field.IsValid() || field.Kind() != reflect.String {
				t.Fatalf("Annotation.%s must be string", name)
			}
			field.SetString(text)
		}
	}
	before := reflect.New(input.Type()).Elem()
	before.Set(reflect.MakeSlice(input.Type(), input.Len(), input.Len()))
	reflect.Copy(before, input)
	result := method.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(style), reflect.ValueOf(iterator), input})
	if !reflect.DeepEqual(before.Interface(), input.Interface()) {
		t.Error("FormatAnnotated modified caller annotations")
	}
	if result[0].IsNil() {
		return nil
	}
	return result[0].Interface().(error)
}

func renderAnnotations(t *testing.T, tokens []chroma.Token, annotations []sourceAnnotation, options ...html.Option) string {
	t.Helper()
	var output bytes.Buffer
	if err := formatAnnotated(t, html.New(options...), &output, styles.Get("monokai"), slices.Values(tokens), annotations); err != nil {
		t.Fatal(err)
	}
	return output.String()
}

func assertAnnotationCoverage(t *testing.T, doc annotationDocument, source string, annotations []sourceAnnotation) {
	t.Helper()
	expected := []annotationCharacter{}
	for offset, character := range source {
		ids := []string{}
		for _, annotation := range annotations {
			if offset >= annotation.Start && offset < annotation.End {
				ids = append(ids, annotation.ID)
			}
		}
		if len(ids) != 0 {
			expected = append(expected, annotationCharacter{character, ids})
		}
	}
	if !reflect.DeepEqual(doc.covered, expected) {
		t.Fatalf("coverage:\ngot %#v\nwant %#v", doc.covered, expected)
	}
	for _, annotation := range annotations {
		if doc.marked[annotation.ID] != source[annotation.Start:annotation.End] {
			t.Fatalf("incorrect text for %q: %q", annotation.ID, doc.marked[annotation.ID])
		}
	}
}

func TestAnnotationsCrossingTokensLinesAndUnicode(t *testing.T) {
	tokens := []chroma.Token{{Type: chroma.Keyword, Value: "ab<&"}, {Type: chroma.Name, Value: "é\n"}, {Type: chroma.Text, Value: "cd🙂z"}}
	source := "ab<&é\ncd🙂z"
	annotations := []sourceAnnotation{{Start: 2, End: 13, ID: "outer"}, {Start: 0, End: 7, ID: "first"}, {Start: 4, End: 9, ID: "middle"}, {Start: 7, End: 13, ID: "last"}}
	output := renderAnnotations(t, tokens, annotations, html.PreventSurroundingPre(true))
	doc := inspectAnnotations(t, output)
	assertAnnotationCoverage(t, doc, source, annotations)
	if doc.text != source {
		t.Fatalf("source text changed: %q", doc.text)
	}
}

func TestAnnotationsAdjacentIdenticalAndNestedRanges(t *testing.T) {
	source := "0123456789"
	annotations := []sourceAnnotation{{0, 5, "a", "", ""}, {5, 10, "b", "", ""}, {0, 5, "c", "", ""}, {2, 4, "d", "", ""}}
	doc := inspectAnnotations(t, renderAnnotations(t, []chroma.Token{{Type: chroma.Text, Value: source}}, annotations, html.PreventSurroundingPre(true)))
	assertAnnotationCoverage(t, doc, source, annotations)
	if doc.text != source {
		t.Fatal(doc.text)
	}
}

type annotationPreWrapper struct{}

func (annotationPreWrapper) Start(code bool, style string) string { return "<section>" }
func (annotationPreWrapper) End(code bool) string                 { return "</section>" }

func TestAnnotationsComposeWithLayouts(t *testing.T) {
	source := "if x\nthen y\n"
	annotations := []sourceAnnotation{{Start: 1, End: len(source) - 1, ID: "range", Class: "diagnostic", Title: "message"}}
	for index, options := range [][]html.Option{
		{html.WithClasses(true), html.ClassPrefix("syntax-")},
		{html.WithClasses(false)},
		{html.Standalone(true), html.WithClasses(true)},
		{html.InlineCode(true)},
		{html.PreventSurroundingPre(true)},
		{html.WithPreWrapper(annotationPreWrapper{})},
		{html.WithLineNumbers(true), html.BaseLineNumber(41)},
		{html.WithLineNumbers(true), html.LineNumbersInTable(true), html.WithClasses(true)},
		{html.WithLinePrompts("$ ", [][2]int{{1, 2}}), html.HighlightLines([][2]int{{1, 2}})},
	} {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			output := renderAnnotations(t, []chroma.Token{{Type: chroma.Keyword, Value: source}}, annotations, options...)
			doc := inspectAnnotations(t, output)
			assertAnnotationCoverage(t, doc, source, annotations)
			if doc.attributes["range"]["class"] != "diagnostic" {
				t.Fatal("annotation class prefixed or lost")
			}
			if doc.attributes["range"]["title"] != "message" {
				t.Fatal("annotation title lost")
			}
			if index == 0 && !strings.Contains(output, `class="syntax-k"`) {
				t.Fatal("syntax styling lost")
			}
			if index == 1 && !strings.Contains(output, `style="color:`) {
				t.Fatal("inline syntax styling lost")
			}
		})
	}
}

func TestAnnotationsEscapeMetadataAndSource(t *testing.T) {
	source := `<script>&"'`
	annotation := sourceAnnotation{0, len(source), `id"<&`, `a" onclick="bad<&`, `title"<&'`}
	output := renderAnnotations(t, []chroma.Token{{Type: chroma.Text, Value: source}}, []sourceAnnotation{annotation}, html.PreventSurroundingPre(true))
	doc := inspectAnnotations(t, output)
	if doc.text != source {
		t.Fatalf("source changed: %q", doc.text)
	}
	attrs := doc.attributes[annotation.ID]
	if attrs["data-annotation"] != annotation.ID || attrs["class"] != annotation.Class || attrs["title"] != annotation.Title {
		t.Fatal(attrs)
	}
	if strings.Contains(output, "<script>") {
		t.Fatal("unescaped source")
	}
}

func TestAnnotationsInvalidInputWritesNothing(t *testing.T) {
	for index, test := range []struct {
		source      string
		annotations []sourceAnnotation
	}{
		{"abc", []sourceAnnotation{{-1, 1, "a", "", ""}}},
		{"abc", []sourceAnnotation{{1, 1, "a", "", ""}}},
		{"abc", []sourceAnnotation{{2, 1, "a", "", ""}}},
		{"abc", []sourceAnnotation{{0, 4, "a", "", ""}}},
		{"abc", []sourceAnnotation{{0, 1, "", "", ""}}},
		{"abc", []sourceAnnotation{{0, 1, "a", "", ""}, {2, 3, "a", "", ""}}},
		{"éx", []sourceAnnotation{{1, 3, "a", "", ""}}},
		{"éx", []sourceAnnotation{{0, 1, "a", "", ""}}},
		{"", []sourceAnnotation{{0, 1, "a", "", ""}}},
		{"\xff", []sourceAnnotation{{0, 1, "a", "", ""}}},
	} {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			var output bytes.Buffer
			f := html.New(html.Standalone(true))
			err := formatAnnotated(t, f, &output, styles.Fallback, slices.Values([]chroma.Token{{Type: chroma.Text, Value: test.source}}), test.annotations)
			if err == nil || output.Len() != 0 {
				t.Fatalf("error=%v output=%q", err, output.String())
			}
		})
	}
}

func TestAnnotationsConcurrentAndSequentialReuse(t *testing.T) {
	annotations := []sourceAnnotation{{0, 2, "id", "class", "title"}}
	f := html.New(html.PreventSurroundingPre(true))
	if !reflect.ValueOf(f).MethodByName("FormatAnnotated").IsValid() {
		t.Fatal("Formatter.FormatAnnotated is missing")
	}

	var group sync.WaitGroup
	errors := make(chan error, 12)
	for i := 0; i < 12; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			var output bytes.Buffer
			calls := 0
			iterator := func(yield func(chroma.Token) bool) { calls++; yield(chroma.Token{Type: chroma.Text, Value: "abc"}) }
			if err := formatAnnotated(t, f, &output, styles.Fallback, iterator, annotations); err != nil {
				errors <- err
				return
			}
			if calls != 1 || !strings.Contains(output.String(), `data-annotation="id"`) || strings.Contains(output.String(), "bad") {
				errors <- fmt.Errorf("configuration or iterator state lost: %s", output.String())
			}
		}()
	}
	group.Wait()
	close(errors)
	for err := range errors {
		t.Error(err)
	}
	var output bytes.Buffer
	if err := formatAnnotated(t, f, &output, styles.Fallback, slices.Values([]chroma.Token{{Type: chroma.Text, Value: "xy"}}), []sourceAnnotation{{1, 2, "id", "", ""}}); err != nil {
		t.Fatal(err)
	}
	if inspectAnnotations(t, output.String()).marked["id"] != "y" {
		t.Fatal("render state retained")
	}
}

type annotationFailWriter struct {
	err   error
	short bool
}

func (w annotationFailWriter) Write(p []byte) (int, error) {
	if w.short {
		return 0, nil
	}
	return 0, w.err
}

func TestAnnotationsPropagateWriterFailures(t *testing.T) {
	sentinel := errors.New("writer failed")
	for _, writer := range []annotationFailWriter{{err: sentinel}, {short: true}} {
		f := html.New()
		err := formatAnnotated(t, f, writer, styles.Fallback, slices.Values([]chroma.Token{{Type: chroma.Text, Value: "x"}}), []sourceAnnotation{{0, 1, "id", "", ""}})
		want := sentinel
		if writer.short {
			want = io.ErrShortWrite
		}
		if !errors.Is(err, want) {
			t.Fatalf("got %v, want %v", err, want)
		}
	}
}

func TestAnnotationsEmptyOptionPreservesExistingOutput(t *testing.T) {
	tokens := []chroma.Token{{Type: chroma.Keyword, Value: "if\nx"}}
	var original, explicit bytes.Buffer
	options := []html.Option{html.WithLineNumbers(true), html.WithClasses(true)}
	if err := html.New(options...).Format(&original, styles.Fallback, slices.Values(tokens)); err != nil {
		t.Fatal(err)
	}

	if err := formatAnnotated(t, html.New(options...), &explicit, styles.Fallback, slices.Values(tokens), nil); err != nil {
		t.Fatal(err)
	}
	if explicit.String() != original.String() {
		t.Fatal("empty annotations changed output")
	}
}
