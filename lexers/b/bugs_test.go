package b_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/b"
)

func TestBugs_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/bugs_basic.bug")
	assert.NoError(t, err)

	analyser, ok := b.Bugs.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.7), analyser.AnalyseText(string(data)))
}
