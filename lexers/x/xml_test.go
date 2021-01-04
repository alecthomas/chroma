package x_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/x"
)

func TestXML_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/xml_doctype_html.xml")
	assert.NoError(t, err)

	analyser, ok := x.XML.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.45), analyser.AnalyseText(string(data)))
}
