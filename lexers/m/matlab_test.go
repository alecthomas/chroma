package m_test

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/m"
)

func TestMatlab_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Text     string
		Expected float32
	}{
		"function": {
			Text: strings.Join([]string{
				"% comment",
				"function foo = bar(a, b, c)",
			}, "\n"),
			Expected: 1.0,
		},
		"comment": {
			Text: strings.Join([]string{
				"",
				"% comment",
				"",
			}, "\n"),
			Expected: 0.2,
		},
		"system cmd": {
			Text: strings.Join([]string{
				"",
				"!rmdir oldtests",
			}, "\n"),
			Expected: 0.2,
		},
		"windows": {
			Text: strings.Join([]string{
				"% comment",
				"function foo = bar(a, b, c)",
			}, "\r\n"),
			Expected: 1.0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			analyser, ok := m.Matlab.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(test.Text))
		})
	}
}
