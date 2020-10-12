package m_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/m"
)

func TestMatlab_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"function": {
			Filepath: "testdata/matlab_function.m",
			Expected: 1.0,
		},
		"comment": {
			Filepath: "testdata/matlab_comment.m",
			Expected: 0.2,
		},
		"systemcmd": {
			Filepath: "testdata/matlab_systemcmd.m",
			Expected: 0.2,
		},
		"windows": {
			Filepath: "testdata/matlab_windows.m",
			Expected: 1.0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := m.Matlab.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
