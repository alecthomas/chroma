package p_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/p"
)

func TestPawn_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/pawn_tagof.pwn")
	assert.NoError(t, err)

	analyser, ok := p.Pawn.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.01), analyser.AnalyseText(string(data)))
}
