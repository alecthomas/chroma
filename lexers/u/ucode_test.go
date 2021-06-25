package u_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/u"
)

func TestUcode_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"endsuspend": {
			Filepath: "testdata/ucode_endsuspend.u",
			Expected: 0.1,
		},
		"endrepeat": {
			Filepath: "testdata/ucode_endrepeat.u",
			Expected: 0.1,
		},
		"variable set": {
			Filepath: "testdata/ucode_varset.u",
			Expected: 0.01,
		},
		"procedure": {
			Filepath: "testdata/ucode_procedure.u",
			Expected: 0.01,
		},
		"self": {
			Filepath: "testdata/ucode_self.u",
			Expected: 0.5,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := u.Ucode.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
