package s_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/s"
)

func TestMatlab_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"only header": {
			Filepath: "testdata/singularity_only_header.def",
			Expected: 0.5,
		},
		"only section": {
			Filepath: "testdata/singularity_only_section.def",
			Expected: 0.49,
		},
		"full": {
			Filepath: "testdata/singularity_full.def",
			Expected: 0.99,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := s.Singularity.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
