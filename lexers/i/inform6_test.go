package i_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/i"
)

func TestInform6_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/inform6_basic.inf")
	assert.NoError(t, err)

	analyser, ok := i.Inform6.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.05), analyser.AnalyseText(string(data)))
}
