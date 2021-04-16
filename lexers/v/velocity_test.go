package v_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/v"
)

func TestVelocity_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"macro": {
			Filepath: "testdata/velocity_macro.vm",
			Expected: 0.26,
		},
		"if": {
			Filepath: "testdata/velocity_if.vm",
			Expected: 0.16,
		},
		"foreach": {
			Filepath: "testdata/velocity_foreach.vm",
			Expected: 0.16,
		},
		"reference": {
			Filepath: "testdata/velocity_reference.vm",
			Expected: 0.01,
		},
		"all": {
			Filepath: "testdata/velocity_all.vm",
			Expected: 0.16,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := v.Velocity.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
