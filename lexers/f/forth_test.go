package f_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/f"
)

func TestForth_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/forth_command.frt")
	assert.NoError(t, err)

	analyser, ok := f.Forth.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.1), analyser.AnalyseText(string(data)))
}
