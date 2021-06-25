package c_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/c"
)

func TestCbmBasicV2_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/cbmbasicv2_basic.bas")
	assert.NoError(t, err)

	analyser, ok := c.CbmBasicV2.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.2), analyser.AnalyseText(string(data)))
}
