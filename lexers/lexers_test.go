package lexers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alecthomas/chroma"
)

// Test source files are in the form <key>.<key> and validation data is in the form <key>.<key>.expected.
func TestLexers(t *testing.T) {
	for _, lexer := range Registry.Lexers {
		name := strings.ToLower(lexer.Config().Name)
		filename := filepath.Join("testdata", name+"."+name)
		expectedFilename := filepath.Join("testdata", name+".expected")
		if _, err := os.Stat(filename); err != nil {
			continue
		}
		if !assert.NotNil(t, lexer) {
			continue
		}
		t.Run(lexer.Config().Name, func(t *testing.T) {
			// Read and tokenise source text.
			actualText, err := ioutil.ReadFile(filename)
			if !assert.NoError(t, err) {
				return
			}
			actual, err := chroma.Tokenise(lexer, nil, string(actualText))
			if !assert.NoError(t, err) {
				return
			}

			// Read expected JSON into token slice.
			expected := []*chroma.Token{}
			r, err := os.Open(expectedFilename)
			if !assert.NoError(t, err) {
				return
			}
			err = json.NewDecoder(r).Decode(&expected)
			if !assert.NoError(t, err) {
				return
			}

			// Equal?
			assert.Equal(t, expected, actual)
		})
	}
}
