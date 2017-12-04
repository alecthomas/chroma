package html

import (
	"errors"
	"io/ioutil"
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
		it, err := lexers.Go.Tokenise(nil, "package main\nfunc main()\n{\nprintln(`hello world`)\n}\n")
		assert.NoError(b, err)
		err = formatter.Format(ioutil.Discard, styles.Fallback, it)
		assert.NoError(b, err)
	}
}

func TestSplitTokensIntoLines(t *testing.T) {
	in := []*chroma.Token{
		{Value: "hello", Type: chroma.NameKeyword},
		{Value: " world\nwhat?\n", Type: chroma.NameKeyword},
	}
	expected := [][]*chroma.Token{
		{
			{Type: chroma.NameKeyword, Value: "hello"},
			{Type: chroma.NameKeyword, Value: " world\n"},
		},
		{
			{Type: chroma.NameKeyword, Value: "what?\n"},
		},
		{
			{Type: chroma.NameKeyword},
		},
	}
	actual := splitTokensIntoLines(in)
	assert.Equal(t, expected, actual)
}

func TestIteratorPanicRecovery(t *testing.T) {
	it := func() *chroma.Token {
		panic(errors.New("bad"))
	}
	err := New().Format(ioutil.Discard, styles.Fallback, it)
	assert.Error(t, err)
}

func Test_formatStyle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{s: ""},
			want: "",
		}, {
			name: "semicolon only",
			args: args{s: ";"},
			want: "",
		}, {
			name: "must strip empty rule at beginning",
			args: args{s: "; margin: 0;"},
			want: "margin: 0;",
		}, {
			name: "must end with semicolon",
			args: args{s: "margin: 0;"},
			want: "margin: 0;",
		}, {
			name: "must remove empty rules in between",
			args: args{s: "margin: 0; ; ; padding: 0;; color: #fff;"},
			want: "margin: 0; padding: 0; color: #fff;",
		}, {
			name: "complex example",
			args: args{s: "   ;margin: 0  ; ; padding: 0;;; ; color: #fff   "},
			want: "margin: 0; padding: 0; color: #fff;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatStyle(tt.args.s); got != tt.want {
				t.Errorf("formatStyle() = %q, want %q", got, tt.want)
			}
		})
	}
}
