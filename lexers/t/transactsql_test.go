package t_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	lt "github.com/alecthomas/chroma/lexers/t"
)

func TestTransactSQL_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"declare": {
			Filepath: "testdata/transactsql_declare.sql",
			Expected: 1.0,
		},
		"bracket": {
			Filepath: "testdata/transactsql_bracket.sql",
			Expected: 0.5,
		},
		"variable": {
			Filepath: "testdata/transactsql_variable.sql",
			Expected: 0.1,
		},
		"go": {
			Filepath: "testdata/transactsql_go.sql",
			Expected: 0.1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := lt.TransactSQL.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
