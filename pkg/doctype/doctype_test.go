package doctype_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma/pkg/doctype"
	"github.com/stretchr/testify/require"
)

func TestDoctype_MatchString(t *testing.T) {
	tests := map[string]struct {
		Text     string
		Pattern  string
		Expected bool
	}{
		"simple html match": {
			Text:     "<!DOCTYPE html> <html>",
			Pattern:  `html.*`,
			Expected: true,
		},
		"full html match": {
			Text:     "<?xml ?><!DOCTYPE html PUBLIC  \"-//W3C//DTD XHTML 1.0 Strict//EN\">",
			Pattern:  `html`,
			Expected: true,
		},
		"missing exclamation mark": {
			Text:     "<?xml ?> <DOCTYPE html PUBLIC \"a\"> <html>",
			Pattern:  `html.*`,
			Expected: false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := doctype.MatchString(test.Text, test.Pattern)
			require.NoError(t, err)

			assert.Equal(t, test.Expected, res)
		})
	}
}
