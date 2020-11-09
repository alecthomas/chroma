package t_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	lt "github.com/alecthomas/chroma/lexers/t"
)

func TestTasm_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/tasm.asm")
	assert.NoError(t, err)

	analyser, ok := lt.Tasm.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
