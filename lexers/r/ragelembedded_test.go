package r_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/r"
)

func TestRagelEmbedded_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/ragel.rl")
	assert.NoError(t, err)

	analyser, ok := r.RagelEmbedded.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
