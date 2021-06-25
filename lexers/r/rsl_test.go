package r_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/r"
)

func TestRSL_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/raise.rsl")
	assert.NoError(t, err)

	analyser, ok := r.RSL.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
