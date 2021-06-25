package p_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/p"
)

func TestProlog_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/prolog.ecl")
	assert.NoError(t, err)

	analyser, ok := p.Prolog.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
