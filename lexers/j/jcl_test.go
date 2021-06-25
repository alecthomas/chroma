package j_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/j"
)

func TestJcl_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/jcl_job_header.jcl")
	assert.NoError(t, err)

	analyser, ok := j.Jcl.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
