package s_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/s"
)

func TestStan_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/stan_basic.stan")
	assert.NoError(t, err)

	analyser, ok := s.Stan.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
