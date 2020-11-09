package c_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/c"
)

func TestCoq_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/coq_reserved_keyword.v")
	assert.NoError(t, err)

	analyser, ok := c.Coq.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
