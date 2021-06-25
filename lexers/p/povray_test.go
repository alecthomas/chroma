package p_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/p"
)

func TestPovRay_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"camera": {
			Filepath: "testdata/povray_camera.pov",
			Expected: 0.05,
		},
		"light_source": {
			Filepath: "testdata/povray_light_source.pov",
			Expected: 0.1,
		},
		"declare": {
			Filepath: "testdata/povray_declare.pov",
			Expected: 0.05,
		},
		"version": {
			Filepath: "testdata/povray_version.pov",
			Expected: 0.05,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := p.Povray.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
