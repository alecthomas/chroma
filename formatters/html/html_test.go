package html

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/alecthomas/assert"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func TestCompressStyle(t *testing.T) {
	style := "color: #888888; background-color: #faffff"
	actual := compressStyle(style)
	expected := "color:#888;background-color:#faffff"
	assert.Equal(t, expected, actual)
}

func BenchmarkHTMLFormatter(b *testing.B) {
	formatter := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it, err := lexers.Get("go").Tokenise(nil, "package main\nfunc main()\n{\nprintln(`hello world`)\n}\n")
		assert.NoError(b, err)
		err = formatter.Format(ioutil.Discard, styles.Fallback, it)
		assert.NoError(b, err)
	}
}

func TestSplitTokensIntoLines(t *testing.T) {
	in := []chroma.Token{
		{Value: "hello", Type: chroma.NameKeyword},
		{Value: " world\nwhat?\n", Type: chroma.NameKeyword},
	}
	expected := [][]chroma.Token{
		{
			{Type: chroma.NameKeyword, Value: "hello"},
			{Type: chroma.NameKeyword, Value: " world\n"},
		},
		{
			{Type: chroma.NameKeyword, Value: "what?\n"},
		},
	}
	actual := chroma.SplitTokensIntoLines(in)
	assert.Equal(t, expected, actual)
}

func TestFormatterStyleToCSS(t *testing.T) {
	builder := styles.Get("github").Builder()
	builder.Add(chroma.LineHighlight, "bg:#ffffcc")
	builder.Add(chroma.LineNumbers, "bold")
	style, err := builder.Build()
	if err != nil {
		t.Error(err)
	}
	formatter := New(WithClasses(true))
	css := formatter.styleToCSS(style)
	for _, s := range css {
		if strings.HasPrefix(strings.TrimSpace(s), ";") {
			t.Errorf("rule starts with semicolon - expected valid css rule without semicolon: %v", s)
		}
	}
}

func TestClassPrefix(t *testing.T) {
	wantPrefix := "some-prefix-"
	withPrefix := New(WithClasses(true), ClassPrefix(wantPrefix))
	noPrefix := New(WithClasses(true))
	for st := range chroma.StandardTypes {
		if noPrefix.class(st) == "" {
			if got := withPrefix.class(st); got != "" {
				t.Errorf("Formatter.class(%v): prefix shouldn't be added to empty classes", st)
			}
		} else if got := withPrefix.class(st); !strings.HasPrefix(got, wantPrefix) {
			t.Errorf("Formatter.class(%v): %q should have a class prefix", st, got)
		}
	}

	var styleBuf bytes.Buffer
	err := withPrefix.WriteCSS(&styleBuf, styles.Fallback)
	assert.NoError(t, err)
	if !strings.Contains(styleBuf.String(), ".some-prefix-chroma ") {
		t.Error("Stylesheets should have a class prefix")
	}
}

func TestTableLineNumberNewlines(t *testing.T) {
	f := New(WithClasses(true), WithLineNumbers(true), LineNumbersInTable(true))
	it, err := lexers.Get("go").Tokenise(nil, "package main\nfunc main()\n{\nprintln(`hello world`)\n}\n")
	assert.NoError(t, err)

	var buf bytes.Buffer
	err = f.Format(&buf, styles.Fallback, it)
	assert.NoError(t, err)

	// Don't bother testing the whole output, just verify it's got line numbers
	// in a <pre>-friendly format.
	// Note: placing the newlines inside the <span> lets browser selections look
	// better, instead of "skipping" over the span margin.
	assert.Contains(t, buf.String(), `<span class="lnt">2
</span><span class="lnt">3
</span><span class="lnt">4
</span>`)
}

func TestTableLineNumberSpacing(t *testing.T) {
	testCases := []struct {
		baseLineNumber int
		expectedBuf    string
	}{{
		7,
		`<span class="lnt"> 7
</span><span class="lnt"> 8
</span><span class="lnt"> 9
</span><span class="lnt">10
</span><span class="lnt">11
</span>`,
	}, {
		6,
		`<span class="lnt"> 6
</span><span class="lnt"> 7
</span><span class="lnt"> 8
</span><span class="lnt"> 9
</span><span class="lnt">10
</span>`,
	}, {
		5,
		`<span class="lnt">5
</span><span class="lnt">6
</span><span class="lnt">7
</span><span class="lnt">8
</span><span class="lnt">9
</span>`,
	}}
	for i, testCase := range testCases {
		f := New(
			WithClasses(true),
			WithLineNumbers(true),
			LineNumbersInTable(true),
			BaseLineNumber(testCase.baseLineNumber),
		)
		it, err := lexers.Get("go").Tokenise(nil, "package main\nfunc main()\n{\nprintln(`hello world`)\n}\n")
		assert.NoError(t, err)
		var buf bytes.Buffer
		err = f.Format(&buf, styles.Fallback, it)
		assert.NoError(t, err, "Test Case %d", i)
		assert.Contains(t, buf.String(), testCase.expectedBuf, "Test Case %d", i)
	}
}

func TestWithPreWrapper(t *testing.T) {
	wrapper := preWrapper{
		start: func(code bool, styleAttr string) string {
			return fmt.Sprintf("<foo%s id=\"code-%t\">", styleAttr, code)
		},
		end: func(code bool) string {
			return fmt.Sprintf("</foo>")
		},
	}

	format := func(f *Formatter) string {
		it, err := lexers.Get("bash").Tokenise(nil, "echo FOO")
		assert.NoError(t, err)

		var buf bytes.Buffer
		err = f.Format(&buf, styles.Fallback, it)
		assert.NoError(t, err)

		return buf.String()
	}

	t.Run("Regular", func(t *testing.T) {
		s := format(New(WithClasses(true)))
		assert.Equal(t, s, `<pre class="chroma"><span class="nb">echo</span> FOO</pre>`)
	})

	t.Run("PreventSurroundingPre", func(t *testing.T) {
		s := format(New(PreventSurroundingPre(true), WithClasses(true)))
		assert.Equal(t, s, `<span class="nb">echo</span> FOO`)
	})

	t.Run("Wrapper", func(t *testing.T) {
		s := format(New(WithPreWrapper(wrapper), WithClasses(true)))
		assert.Equal(t, s, `<foo class="chroma" id="code-true"><span class="nb">echo</span> FOO</foo>`)
	})

	t.Run("Wrapper, LineNumbersInTable", func(t *testing.T) {
		s := format(New(WithPreWrapper(wrapper), WithClasses(true), WithLineNumbers(true), LineNumbersInTable(true)))

		assert.Equal(t, s, `<div class="chroma">
<table class="lntable"><tr><td class="lntd">
<foo class="chroma" id="code-false"><span class="lnt">1
</span></foo></td>
<td class="lntd">
<foo class="chroma" id="code-true"><span class="nb">echo</span> FOO</foo></td></tr></table>
</div>
`)
	})
}

func TestReconfigureOptions(t *testing.T) {
	options := []Option{
		WithClasses(true),
		WithLineNumbers(true),
	}

	options = append(options, WithLineNumbers(false))

	f := New(options...)

	it, err := lexers.Get("bash").Tokenise(nil, "echo FOO")
	assert.NoError(t, err)

	var buf bytes.Buffer
	err = f.Format(&buf, styles.Fallback, it)

	assert.NoError(t, err)
	assert.Equal(t, `<pre class="chroma"><span class="nb">echo</span> FOO</pre>`, buf.String())
}
