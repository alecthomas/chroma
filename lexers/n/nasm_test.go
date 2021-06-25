package n_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/n"
)

func TestNasm_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/nasm.asm")
	assert.NoError(t, err)

	analyser, ok := n.Nasm.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0), analyser.AnalyseText(string(data)))
}
