package b_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/b"
)

func TestBrainfuck_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"plus minus": {
			Filepath: "testdata/brainfuck_plus_minus.bf",
			Expected: 1.0,
		},
		"greater less": {
			Filepath: "testdata/brainfuck_greater_less.bf",
			Expected: 1.0,
		},
		"minus only": {
			Filepath: "testdata/brainfuck_minus.bf",
			Expected: 0.5,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := b.Brainfuck.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
