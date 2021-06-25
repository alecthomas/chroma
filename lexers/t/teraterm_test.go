package t_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	lt "github.com/alecthomas/chroma/lexers/t"
)

func TestTeraTerm_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/teraterm_commands.ttl")
	assert.NoError(t, err)

	analyser, ok := lt.TeraTerm.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.01), analyser.AnalyseText(string(data)))
}
