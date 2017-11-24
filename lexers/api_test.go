package lexers_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
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
		assert.Equal(t, lexers.Get("xml"), lexers.XML)
	})
	t.Run("ByAlias", func(t *testing.T) {
		assert.Equal(t, lexers.Get("as"), lexers.Actionscript)
	})
	t.Run("ViaFilename", func(t *testing.T) {
		assert.Equal(t, lexers.Get("svg"), lexers.XML)
	})
}
