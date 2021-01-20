package s_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/s"
)

func TestSlurm_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/slurm.sl")
	assert.NoError(t, err)

	analyser, ok := s.Slurm.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
