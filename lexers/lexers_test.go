package lexers_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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

func TestGlobs(t *testing.T) {
	filename := "main.go"
	for _, lexer := range lexers.Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.Filenames {
			_, err := filepath.Match(glob, filename)
			require.NoError(t, err)
		}
		for _, glob := range config.AliasFilenames {
			_, err := filepath.Match(glob, filename)
			require.NoError(t, err)
		}
	}
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
		// skip text analysis test files
		if file.Name() == "analysis" {
			continue
		}

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

func FileTestAnalysis(t *testing.T, lexer chroma.Lexer, actualFilepath, expectedFilepath string) {
	t.Helper()
	t.Run(lexer.Config().Name+"/"+actualFilepath, func(t *testing.T) {
		expectedData, err := ioutil.ReadFile(expectedFilepath)
		assert.NoError(t, err)

		analyser, ok := lexer.(chroma.Analyser)
		assert.True(t, ok, "lexer %q does not set analyser", lexer.Config().Name)

		data, err := ioutil.ReadFile(actualFilepath)
		assert.NoError(t, err)

		actual := analyser.AnalyseText(string(data))

		if os.Getenv("RECORD") == "true" {
			// Update the expected file with the generated output of this lexer
			f, err := os.Create(expectedFilepath)
			defer f.Close() // nolint: gosec
			assert.NoError(t, err)

			_, err = f.WriteString(strconv.FormatFloat(float64(actual), 'f', -1, 32))
			assert.NoError(t, err)
		} else {
			expected, err := strconv.ParseFloat(strings.TrimSpace(string(expectedData)), 32)
			assert.NoError(t, err)

			assert.Equal(t, float32(expected), actual)
		}
	})
}

func TestLexersTextAnalyser(t *testing.T) {
	files, err := filepath.Glob("testdata/analysis/*.actual")
	assert.NoError(t, err)

	for _, actualFilepath := range files {
		filename := filepath.Base(actualFilepath)
		baseFilename := strings.TrimSuffix(filename, filepath.Ext(filename))
		lexerName := strings.Split(baseFilename, ".")[0]

		lexer := lexers.Get(lexerName)
		assert.NotNil(t, lexer, "no lexer found for name %q", lexerName)

		expectedFilepath := "testdata/analysis/" + baseFilename + ".expected"

		FileTestAnalysis(t, lexer, actualFilepath, expectedFilepath)
	}
}
