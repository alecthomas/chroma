package m_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/m"
)

func TestModula2_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"pascal flavour": {
			Filepath: "testdata/modula2_pascal.def",
			Expected: 0,
		},
		"pascal flavour with function": {
			Filepath: "testdata/modula2_pascal_function.def",
			Expected: 0,
		},
		"basic": {
			Filepath: "testdata/modula2_basic.def",
			Expected: 0.6,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := m.Modula2.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
