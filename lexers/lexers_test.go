package lexers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alecthomas/chroma"
)

// Test source files are in the form <key>.<key> and validation data is in the form <key>.<key>.expected.
func TestLexers(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	require.NoError(t, err)

	for _, file := range files {
		ext := filepath.Ext(file.Name())[1:]
		if ext != "actual" {
			continue
		}

		lexer := Get(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
		if !assert.NotNil(t, lexer) {
			continue
		}

		filename := filepath.Join("testdata", file.Name())
		expectedFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".expected"

		lexer = chroma.Coalesce(lexer)
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
