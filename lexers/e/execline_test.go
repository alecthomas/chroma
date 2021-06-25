package e_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/e"
)

func TestExecline_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/execline_shebang.exec")
	assert.NoError(t, err)

	analyser, ok := e.Execline.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
