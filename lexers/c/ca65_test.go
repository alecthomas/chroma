package c_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/c"
)

func TestCa65_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/ca65_comment.s")
	assert.NoError(t, err)

	analyser, ok := c.Ca65.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.9), analyser.AnalyseText(string(data)))
}
