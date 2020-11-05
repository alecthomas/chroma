package r_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/r"
)

func TestR_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/r_expression.r")
	assert.NoError(t, err)

	analyser, ok := r.R.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.11), analyser.AnalyseText(string(data)))
}
