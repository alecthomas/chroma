package p_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/p"
)

func TestPerl6_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"shebang": {
			Filepath: "testdata/perl6_shebang.pl6",
			Expected: 1.0,
		},
		"v6": {
			Filepath: "testdata/perl6_v6.pl6",
			Expected: 1.0,
		},
		"enum": {
			Filepath: "testdata/perl6_enum.pl6",
			Expected: 0.05,
		},
		"scoped class": {
			Filepath: "testdata/perl6_scoped_class.pl6",
			Expected: 1.0,
		},
		"assignment": {
			Filepath: "testdata/perl6_assign.pl6",
			Expected: 0.4,
		},
		"strip pod": {
			Filepath: "testdata/perl6_pod.pl6",
			Expected: 0.4,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := p.Perl6.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}
