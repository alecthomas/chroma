package html

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

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
	writer, err := formatter.Format(ioutil.Discard, styles.Fallback)
	assert.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = lexers.Go.Tokenise(nil, "package main\nfunc main()\n{\nprintln(`hello world`)\n}\n", writer)
		assert.NoError(b, err)
	}
}
