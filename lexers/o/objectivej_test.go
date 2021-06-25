package o_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/o"
)

func TestObjectiveJ_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/objectivej_import.j")
	assert.NoError(t, err)

	analyser, ok := o.ObjectiveJ.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
