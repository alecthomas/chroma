package s_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/s"
)

func TestSmali_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"class": {
			Filepath: "testdata/smali_class.smali",
			Expected: 0.5,
		},
		"class with keyword": {
			Filepath: "testdata/smali_class_keyword.smali",
			Expected: 0.8,
		},
		"keyword": {
			Filepath: "testdata/smali_keyword.smali",
			Expected: 0.6,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := s.Smali.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
