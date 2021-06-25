package e_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/e"
)

func TestEzhil_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/ezhil_basic.n")
	assert.NoError(t, err)

	analyser, ok := e.Ezhil.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.25), analyser.AnalyseText(string(data)))
}
