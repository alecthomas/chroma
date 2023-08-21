package lexers_test

import (
	"os"
	"testing"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"

	"github.com/alecthomas/assert/v2"
)

func TestActionscript3_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"basic": {
			Filepath: "testdata/actionscript3.as",
			Expected: 0.3,
		},
		"capital letters": {
			Filepath: "testdata/actionscript3_capital_letter.as",
			Expected: 0.3,
		},
		"spaces": {
			Filepath: "testdata/actionscript3_spaces.as",
			Expected: 0.3,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := lexers.Actionscript3.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
