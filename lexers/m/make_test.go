package m_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/m"
)

func TestMakefile_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/makefile")
	assert.NoError(t, err)

	analyser, ok := m.Makefile.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.1), analyser.AnalyseText(string(data)))
}
