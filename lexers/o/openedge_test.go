package o_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/o"
)

func TestOpenEdge_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"end": {
			Filepath: "testdata/openedge_end.p",
			Expected: 0.05,
		},
		"end procedure": {
			Filepath: "testdata/openedge_end_procedure.p",
			Expected: 0.05,
		},
		"else do": {
			Filepath: "testdata/openedge_else_do.p",
			Expected: 0.05,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := o.OpenEdgeABL.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
