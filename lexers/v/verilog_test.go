package v_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/v"
)

func TestVerilog_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"reg": {
			Filepath: "testdata/verilog_reg.v",
			Expected: 0.1,
		},
		"wire": {
			Filepath: "testdata/verilog_wire.v",
			Expected: 0.1,
		},
		"assign": {
			Filepath: "testdata/verilog_assign.v",
			Expected: 0.1,
		},
		"all": {
			Filepath: "testdata/verilog_all.v",
			Expected: 0.3,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := v.Verilog.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
