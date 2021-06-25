package l_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/l"
)

func TestLogos_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/logos_basic.xm")
	assert.NoError(t, err)

	analyser, ok := l.Logos.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
