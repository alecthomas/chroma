package n_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/n"
)

func TestNermerle_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/nemerle_if.n")
	assert.NoError(t, err)

	analyser, ok := n.Nemerle.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.1), analyser.AnalyseText(string(data)))
}
