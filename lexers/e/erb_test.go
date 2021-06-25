package e_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/e"
)

func TestErb_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/erb_basic.erb")
	assert.NoError(t, err)

	analyser, ok := e.Erb.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.4), analyser.AnalyseText(string(data)))
}
