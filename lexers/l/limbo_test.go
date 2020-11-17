package l_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/l"
)

func TestLimbo_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/limbo_basic.b")
	assert.NoError(t, err)

	analyser, ok := l.Limbo.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.7), analyser.AnalyseText(string(data)))
}
