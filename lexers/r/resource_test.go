package r_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/r"
)

func TestResource_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/resource.txt")
	assert.NoError(t, err)

	analyser, ok := r.Resource.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
