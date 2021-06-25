package t_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	lt "github.com/alecthomas/chroma/lexers/t"
)

func TestTads3_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"GameMainDef": {
			Filepath: "testdata/tads3_game_main_def.t",
			Expected: 0.2,
		},
		"__TADS keyword": {
			Filepath: "testdata/tads3_tads_keyword.t",
			Expected: 0.2,
		},
		"version info": {
			Filepath: "testdata/tads3_version_info.t",
			Expected: 0.1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := lt.Tads3.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
