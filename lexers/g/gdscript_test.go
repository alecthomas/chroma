package g_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/g"
)

func TestGdSript_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"func": {
			Filepath: "testdata/gdscript_func.gd",
			Expected: 0.8,
		},
		"keyword first group": {
			Filepath: "testdata/gdscript_keyword.gd",
			Expected: 0.4,
		},
		"keyword second group": {
			Filepath: "testdata/gdscript_keyword2.gd",
			Expected: 0.2,
		},
		"full": {
			Filepath: "testdata/gdscript_full.gd",
			Expected: 1.0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := g.GDScript.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
