package xml_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma/pkg/xml"
)

func TestXML_MatchString(t *testing.T) {
	tests := map[string]struct {
		Text     string
		Pattern  string
		Expected bool
	}{
		"simple xml match": {
			Text:     "<?xml ?><!DOCTYPE html PUBLIC  \"-//W3C//DTD XHTML 1.0 Strict//EN\">",
			Expected: true,
		},
		"xmlns": {
			Text:     "<html xmlns>abc</html>",
			Expected: true,
		},
		"html": {
			Text:     "<html>",
			Expected: false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res := xml.MatchString(test.Text)
			assert.Equal(t, test.Expected, res)
		})
	}
}
