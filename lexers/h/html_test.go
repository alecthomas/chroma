package h_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/h"
)

func TestHTML_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/html_doctype.html")
	assert.NoError(t, err)

	analyser, ok := h.HTML.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.5), analyser.AnalyseText(string(data)))
}
