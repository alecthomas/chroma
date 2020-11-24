package h_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/h"
)

func TestHybris_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"public method": {
			Filepath: "testdata/hybris_public.hyb",
			Expected: 0.01,
		},
		"private method": {
			Filepath: "testdata/hybris_private.hyb",
			Expected: 0.01,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := h.Hybris.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
