package l_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/l"
)

func TestLasso_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"shebang": {
			Filepath: "testdata/lasso_shebang.lasso",
			Expected: 0.8,
		},
		"delimiter": {
			Filepath: "testdata/lasso_delimiter.lasso",
			Expected: 0.4,
		},
		"local": {
			Filepath: "testdata/lasso_local.lasso",
			Expected: 0.4,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := l.Lasso.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
