package lexers_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/lexers/a"
	"github.com/alecthomas/chroma/lexers/x"
	"github.com/alecthomas/chroma/styles"
)

func TestCompileAllRegexes(t *testing.T) {
	for _, lexer := range lexers.Registry.Lexers {
		it, err := lexer.Tokenise(nil, "")
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
		err = formatters.NoOp.Format(ioutil.Discard, styles.SwapOff, it)
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
	}
}

func TestGet(t *testing.T) {
	t.Run("ByName", func(t *testing.T) {
		assert.Equal(t, lexers.Get("xml"), x.XML)
	})
	t.Run("ByAlias", func(t *testing.T) {
		assert.Equal(t, lexers.Get("as"), a.Actionscript)
	})
	t.Run("ViaFilename", func(t *testing.T) {
		assert.Equal(t, lexers.Get("svg"), x.XML)
	})
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexers.Get("go")
	}
}

func FileTest(t *testing.T, lexer chroma.Lexer, actualFilename, expectedFilename string) {
	t.Helper()
	t.Run(lexer.Config().Name+"/"+actualFilename, func(t *testing.T) {
		// Read and tokenise source text.
		actualText, err := ioutil.ReadFile(actualFilename)
		assert.NoError(t, err)
		actual, err := chroma.Tokenise(lexer, nil, string(actualText))
		assert.NoError(t, err)

		if os.Getenv("RECORD") == "true" {
			// Update the expected file with the generated output of this lexer
			f, err := os.Create(expectedFilename)
			defer f.Close() // nolint: gosec
			assert.NoError(t, err)
			assert.NoError(t, formatters.JSON.Format(f, nil, chroma.Literator(actual...)))
		} else {
			// Read expected JSON into token slice.
			var expected []chroma.Token
			r, err := os.Open(expectedFilename)
			assert.NoError(t, err)
			err = json.NewDecoder(r).Decode(&expected)
			assert.NoError(t, err)

			// Equal?
			assert.Equal(t, expected, actual)
		}
	})
}

// Test source files are in the form <key>.<key> and validation data is in the form <key>.<key>.expected.
func TestLexers(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	assert.NoError(t, err)

	for _, file := range files {
		if file.IsDir() {
			dirname := filepath.Join("testdata", file.Name())
			lexer := lexers.Get(file.Name())
			assert.NotNil(t, lexer)

			subFiles, err := ioutil.ReadDir(dirname)
			assert.NoError(t, err)

			for _, subFile := range subFiles {
				ext := filepath.Ext(subFile.Name())[1:]
				if ext != "actual" {
					continue
				}

				filename := filepath.Join(dirname, subFile.Name())
				expectedFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".expected"

				lexer = chroma.Coalesce(lexer)
				FileTest(t, lexer, filename, expectedFilename)
			}
		} else {
			ext := filepath.Ext(file.Name())[1:]
			if ext != "actual" {
				continue
			}

			base := strings.Split(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())), ".")[0]
			lexer := lexers.Get(base)
			assert.NotNil(t, lexer)

			filename := filepath.Join("testdata", file.Name())
			expectedFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".expected"

			lexer = chroma.Coalesce(lexer)
			FileTest(t, lexer, filename, expectedFilename)
		}
	}
}
