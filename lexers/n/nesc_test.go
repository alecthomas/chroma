package n_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/n"
)

func TestNesc_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"include": {
			Filepath: "testdata/nesc_include.nc",
			Expected: 0.1,
		},
		"ifdef": {
			Filepath: "testdata/nesc_ifdef.nc",
			Expected: 0.1,
		},
		"ifndef": {
			Filepath: "testdata/nesc_ifndef.nc",
			Expected: 0.1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := n.NesC.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
