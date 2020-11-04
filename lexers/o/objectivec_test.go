package o_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/o"
)

func TestObjectiveC_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"keyword_end": {
			Filepath: "testdata/objectivec_keyword_end.m",
			Expected: 1.0,
		},
		"keyword_implementation": {
			Filepath: "testdata/objectivec_keyword_implementation.m",
			Expected: 1.0,
		},
		"keyword_protocol": {
			Filepath: "testdata/objectivec_keyword_protocol.m",
			Expected: 1.0,
		},
		"nsstring": {
			Filepath: "testdata/objectivec_nsstring.m",
			Expected: 0.8,
		},
		"nsnumber": {
			Filepath: "testdata/objectivec_nsnumber.m",
			Expected: 0.7,
		},
		"message": {
			Filepath: "testdata/objectivec_message.m",
			Expected: 0.8,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := o.ObjectiveC.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
