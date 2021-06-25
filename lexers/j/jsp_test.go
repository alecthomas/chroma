package j_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/j"
)

func TestJsp_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/jsp_basic.jsp")
	assert.NoError(t, err)

	analyser, ok := j.Jsp.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.49), analyser.AnalyseText(string(data)))
}
