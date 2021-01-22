package n_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/n"
)

func TestNumPy_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"basic": {
			Filepath: "testdata/numpy_basic",
			Expected: 1.0,
		},
		"from numpy import": {
			Filepath: "testdata/numpy_from_import",
			Expected: 1.0,
		},
		"regular python": {
			Filepath: "testdata/numpy.py",
			Expected: 1.0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := n.NumPy.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
