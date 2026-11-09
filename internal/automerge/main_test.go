package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
)

const baseLexer = `<lexer>
  <config>
    <name>Real</name>
    <alias>real</alias>
  </config>
  <rules>
    <state name="root">
      <rule pattern="x">
        <token type="Text"/>
        <push state="a" state="b"/>
      </rule>
    </state>
  </rules>
</lexer>`

const baseStyle = `<style name="test" counterpart="test-dark"><entry type="Text" style="#000"/></style>`

func TestFileReason(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		path     string
		head     string
		expected string
	}{
		{
			name:   "TestdataAdded",
			status: "added",
			path:   "lexers/testdata/real/case.actual",
		},
		{
			name:     "TestdataRemoved",
			status:   "removed",
			path:     "lexers/testdata/real/case.actual",
			expected: "lexers/testdata/real/case.actual is removed",
		},
		{
			name:   "LexerRulesChanged",
			status: "modified",
			path:   "lexers/embedded/real.xml",
			head:   strings.Replace(baseLexer, `pattern="x"`, `pattern="y"`, 1),
		},
		{
			name:     "LexerConfigChanged",
			status:   "modified",
			path:     "lexers/embedded/real.xml",
			head:     strings.Replace(baseLexer, "<alias>real</alias>", "<alias>python</alias>", 1),
			expected: "lexers/embedded/real.xml changes lexer config",
		},
		{
			name:   "LexerConfigChangedBehindCommentedDecoy",
			status: "modified",
			path:   "lexers/embedded/real.xml",
			head: strings.Replace(
				strings.Replace(baseLexer, "<alias>real</alias>", "<alias>python</alias>", 1),
				"<lexer>", "<lexer><!-- <config><name>Real</name><alias>real</alias></config> -->", 1),
			expected: "lexers/embedded/real.xml changes lexer config",
		},
		{
			name:     "LexerAdded",
			status:   "added",
			path:     "lexers/embedded/real.xml",
			expected: "lexers/embedded/real.xml is added",
		},
		{
			name:   "StyleEntryChanged",
			status: "modified",
			path:   "styles/test.xml",
			head:   strings.Replace(baseStyle, "#000", "#111", 1),
		},
		{
			name:     "StyleCounterpartChanged",
			status:   "modified",
			path:     "styles/test.xml",
			head:     strings.Replace(baseStyle, "test-dark", "other", 1),
			expected: "styles/test.xml changes style attributes",
		},
		{
			name:     "OtherFile",
			status:   "modified",
			path:     "lexer.go",
			expected: "lexer.go is not a lexer, style or testdata file",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pr := &pullRequest{mergeBase: "base", head: "head", content: func(_ context.Context, path, ref string) (string, error) {
				switch {
				case ref == "head":
					return test.head, nil
				case lexerPath.MatchString(path):
					return baseLexer, nil
				case stylePath.MatchString(path):
					return baseStyle, nil
				default:
					return "", fmt.Errorf("unexpected fetch of %s at %s", path, ref)
				}
			}}
			actual, err := pr.fileReason(t.Context(), test.status, test.path)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
