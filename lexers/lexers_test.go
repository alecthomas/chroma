package lexers_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

func TestCompileAllRegexes(t *testing.T) {
	for _, lexer := range lexers.GlobalLexerRegistry.Lexers {
		it, err := lexer.Tokenise(nil, "")
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
		err = formatters.NoOp.Format(io.Discard, styles.Get("swapoff"), it)
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
	}
}

func TestGet(t *testing.T) {
	t.Run("ByName", func(t *testing.T) {
		assert.True(t, lexers.Get("xml") == lexers.GlobalLexerRegistry.Get("XML"))
	})
	t.Run("ByAlias", func(t *testing.T) {
		assert.True(t, lexers.Get("as") == lexers.GlobalLexerRegistry.Get("Actionscript"))
	})
	t.Run("ViaFilename", func(t *testing.T) {
		expected := lexers.Get("XML")
		actual := lexers.GlobalLexerRegistry.Get("test.svg")
		assert.Equal(t,
			repr.String(expected.Config(), repr.Indent("  ")),
			repr.String(actual.Config(), repr.Indent("  ")))
	})
}

func TestAliases(t *testing.T) {
	t.Run("UseNameIfNoAliases", func(t *testing.T) {
		expected := lexers.GlobalLexerRegistry.Aliases(false)
		actual := lexers.Aliases(false)
		assert.Equal(t, expected, actual)
	})
	t.Run("SkipIfNoAliases", func(t *testing.T) {
		expected := lexers.GlobalLexerRegistry.Aliases(true)
		actual := lexers.Aliases(true)
		assert.Equal(t, expected, actual)
	})
}

func TestGlobs(t *testing.T) {
	filename := "main.go"
	for _, lexer := range lexers.GlobalLexerRegistry.Lexers {
		config := lexer.Config()
		for _, glob := range config.Filenames {
			_, err := filepath.Match(glob, filename)
			assert.NoError(t, err)
		}
		for _, glob := range config.AliasFilenames {
			_, err := filepath.Match(glob, filename)
			assert.NoError(t, err)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	for range b.N {
		lexers.Get("go")
	}
}

func FileTest(t *testing.T, lexer chroma.Lexer, sourceFile, expectedFilename string) {
	t.Helper()
	t.Run(lexer.Config().Name+"/"+sourceFile, func(t *testing.T) {
		// Read and tokenise source text.
		sourceBytes, err := os.ReadFile(sourceFile)
		assert.NoError(t, err)
		actualTokens, err := chroma.Tokenise(lexer, nil, string(sourceBytes))
		assert.NoError(t, err)

		// Check for error tokens early
		for _, token := range actualTokens {
			if token.Type == chroma.Error {
				t.Logf("Found Error token in lexer %s output: %s", lexer.Config().Name, repr.String(token))
			}
		}

		// Use a bytes.Buffer to "render" the actual bytes
		var actualBytes bytes.Buffer
		err = formatters.JSON.Format(&actualBytes, nil, chroma.Literator(actualTokens...))
		assert.NoError(t, err)

		expectedBytes, err := os.ReadFile(expectedFilename)
		assert.NoError(t, err)

		// Check that the expected bytes are identical
		if !bytes.Equal(actualBytes.Bytes(), expectedBytes) {
			if os.Getenv("RECORD") == "true" {
				f, err := os.Create(expectedFilename)
				assert.NoError(t, err)
				_, err = f.Write(actualBytes.Bytes())
				assert.NoError(t, err)
				assert.NoError(t, f.Close())
			} else {
				// fail via an assertion of string values for diff output
				assert.Equal(t, string(expectedBytes), actualBytes.String())
			}
		}
	})
}

// Test source files are in the form <key>.<key> and validation data is in the form <key>.<key>.expected.
func TestLexers(t *testing.T) {
	files, err := os.ReadDir("testdata")
	assert.NoError(t, err)

	for _, file := range files {
		// skip text analysis test files
		if file.Name() == "analysis" {
			continue
		}

		if file.IsDir() {
			dirname := filepath.Join("testdata", file.Name())
			lexer := lexers.Get(file.Name())
			assert.NotZero(t, lexer)

			subFiles, err := os.ReadDir(dirname)
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
			assert.NotZero(t, lexer, base)

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
		expectedData, err := os.ReadFile(expectedFilepath)
		assert.NoError(t, err)

		analyser, ok := lexer.(chroma.Analyser)
		assert.True(t, ok, "lexer %q does not set analyser", lexer.Config().Name)

		data, err := os.ReadFile(actualFilepath)
		assert.NoError(t, err)

		actual := analyser.AnalyseText(string(data))
		var actualData bytes.Buffer
		fmt.Fprintf(&actualData, "%s\n", strconv.FormatFloat(float64(actual), 'f', -1, 32))

		if !bytes.Equal(expectedData, actualData.Bytes()) {
			if os.Getenv("RECORD") == "true" {
				f, err := os.Create(expectedFilepath)
				assert.NoError(t, err)
				_, err = f.Write(actualData.Bytes())
				assert.NoError(t, err)
				assert.NoError(t, f.Close())
			} else {
				assert.Equal(t, string(expectedData), actualData.String())
			}
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
		assert.NotZero(t, lexer, "no lexer found for name %q", lexerName)

		expectedFilepath := "testdata/analysis/" + baseFilename + ".expected"

		FileTestAnalysis(t, lexer, actualFilepath, expectedFilepath)
	}
}
