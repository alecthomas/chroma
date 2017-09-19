package lexers_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func TestCompileAllRegexes(t *testing.T) {
	writer, err := formatters.NoOp.Format(ioutil.Discard, styles.SwapOff)
	assert.NoError(t, err)
	for _, lexer := range lexers.Registry.Lexers {
		err = lexer.Tokenise(nil, "", writer)
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
	}
}
