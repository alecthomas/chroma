package x_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/x"
)

func TestXSLT_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/xslt.xsl")
	assert.NoError(t, err)

	analyser, ok := x.XSLT.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.8), analyser.AnalyseText(string(data)))
}
