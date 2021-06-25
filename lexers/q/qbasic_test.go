package q_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/q"
)

func TestQBasic_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"dynamic_cmd": {
			Filepath: "testdata/qbasic_dynamiccmd.bas",
			Expected: 0.9,
		},
		"static_cmd": {
			Filepath: "testdata/qbasic_staticcmd.bas",
			Expected: 0.9,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := q.Qbasic.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
