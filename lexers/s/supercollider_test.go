package s_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/s"
)

func TestSuperCollider_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"sinosc": {
			Filepath: "testdata/supercollider_sinosc.sc",
			Expected: 0.1,
		},
		"thisFunctionDef": {
			Filepath: "testdata/supercollider_thisfunctiondef.sc",
			Expected: 0.1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := s.SuperCollider.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
