package v_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/v"
)

func TestMatlab_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"module": {
			Filepath: "testdata/vb_module.vb",
			Expected: 0.5,
		},
		"namespace": {
			Filepath: "testdata/vb_namespace.vb",
			Expected: 0.5,
		},
		"if": {
			Filepath: "testdata/vb_if.vb",
			Expected: 0.5,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := v.VBNet.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
