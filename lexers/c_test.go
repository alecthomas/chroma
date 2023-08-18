package lexers_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"

	"github.com/alecthomas/assert/v2"
)

func TestC_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"include": {
			Filepath: "testdata/c_include.c",
			Expected: 0.1,
		},
		"ifdef": {
			Filepath: "testdata/c_ifdef.c",
			Expected: 0.1,
		},
		"ifndef": {
			Filepath: "testdata/c_ifndef.c",
			Expected: 0.1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := lexers.C.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
