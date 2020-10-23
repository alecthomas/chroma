package f_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/f"
)

func TestFsharp_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"pipeline operator": {
			Filepath: "testdata/fsharp_pipeline_operator.fs",
			Expected: 0.1,
		},
		"forward pipeline operator": {
			Filepath: "testdata/fsharp_forward_pipeline_operator.fs",
			Expected: 0.05,
		},
		"backward pipeline operator": {
			Filepath: "testdata/fsharp_backward_pipeline_operator.fs",
			Expected: 0.05,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := f.Fsharp.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
