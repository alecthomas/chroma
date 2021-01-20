package s_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/s"
)

func TestSSP_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/ssp_basic.ssp")
	assert.NoError(t, err)

	analyser, ok := s.SSP.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.9), analyser.AnalyseText(string(data)))
}
