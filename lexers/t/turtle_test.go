package t_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	lt "github.com/alecthomas/chroma/lexers/t"
)

func TestMatlab_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/turtle_basic.ttl")
	assert.NoError(t, err)

	analyser, ok := lt.Turtle.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.8), analyser.AnalyseText(string(data)))
}
