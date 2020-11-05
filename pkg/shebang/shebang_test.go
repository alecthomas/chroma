package shebang_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma/pkg/shebang"
	"github.com/stretchr/testify/require"
)

func TestShebang_MatchString(t *testing.T) {
	tests := map[string]struct {
		Text     string
		Pattern  string
		Expected bool
	}{
		"with break line": {
			Text:     "#!/usr/bin/env python\n",
			Pattern:  `python(2\.\d)?`,
			Expected: true,
		},
		"full match": {
			Text:     "#!/usr/bin/python2.4",
			Pattern:  `python(2\.\d)?`,
			Expected: true,
		},
		"start something with": {
			Text:     "#!/usr/bin/startsomethingwith python",
			Pattern:  `python(2\.\d)?`,
			Expected: true,
		},
		"windows path": {
			Text:     "#!C:\\Python2.4\\Python.exe",
			Pattern:  `python(2\.\d)?`,
			Expected: true,
		},
		"path with dash": {
			Text:     "#!/usr/bin/python-ruby",
			Pattern:  `python(2\.\d)?`,
			Expected: false,
		},
		"path with slash": {
			Text:     "#!/usr/bin/python/ruby",
			Pattern:  `python(2\.\d)?`,
			Expected: false,
		},
		"only shebang": {
			Text:     "#!",
			Pattern:  `python`,
			Expected: false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := shebang.MatchString(test.Text, test.Pattern)
			require.NoError(t, err)

			assert.Equal(t, test.Expected, res)
		})
	}
}
