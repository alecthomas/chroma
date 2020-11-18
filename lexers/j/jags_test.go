package j_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/j"
)

func TestJags_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"model only": {
			Filepath: "testdata/jags_model.jag",
			Expected: 0.3,
		},
		"model and data": {
			Filepath: "testdata/jags_data.jag",
			Expected: 0.9,
		},
		"model and var": {
			Filepath: "testdata/jags_var.jag",
			Expected: 0.9,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := j.Jags.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
