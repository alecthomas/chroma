//go:build chroma_no_lexers

package lexers_test

import (
	"bytes"
	"slices"
	"testing"
	"testing/fstest"

	assert "github.com/alecthomas/assert/v2"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
	"github.com/alecthomas/chroma/v3/quick"
)

func TestWithoutBundledLexers(t *testing.T) {
	assert.Equal(t, []string{}, lexers.Names(false))
	assert.Equal(t, []string{}, lexers.Aliases(false))
	assert.Zero(t, lexers.Get("go"))
	assert.Zero(t, lexers.Match("main.go"))
	assert.Zero(t, lexers.MatchMimeType("text/x-gosrc"))
	assert.Zero(t, lexers.Analyse("package main\nfunc main() {}\n"))

	registry := lexers.GlobalLexerRegistry
	lexers.GlobalLexerRegistry = chroma.NewLexerRegistry()
	t.Cleanup(func() { lexers.GlobalLexerRegistry = registry })
	definitions := fstest.MapFS{
		"words.xml": {Data: []byte(`<lexer><config><name>Local Words</name><alias>local-words</alias></config>
<rules><state name="root">
<rule pattern="[a-z]+"><token type="Name"/></rule>
<rule pattern="\s+"><token type="Text"/></rule>
</state></rules></lexer>`)},
		"outer.xml": {Data: []byte(`<lexer><config><name>Local Outer</name><alias>local-outer</alias>
<filename>*.local</filename><mime_type>text/x-local</mime_type></config>
<rules><state name="root"><rule pattern=".+"><using lexer="local-words"/></rule></state></rules></lexer>`)},
	}
	for _, name := range []string{"words.xml", "outer.xml"} {
		lexer, err := chroma.NewXMLLexer(definitions, name)
		assert.NoError(t, err)
		lexers.Register(lexer)
	}
	assert.Equal(t, []string{"Local Outer", "Local Words"}, lexers.Names(false))
	lexer := lexers.Get("local-outer")
	assert.Equal(t, lexer, lexers.Match("example.local"))
	assert.Equal(t, lexer, lexers.MatchMimeType("text/x-local"))
	tokens, err := chroma.Tokenise(lexer, nil, "hello world")
	assert.NoError(t, err)
	assert.Equal(t, []chroma.Token{{Type: chroma.Name, Value: "hello"}, {Type: chroma.Text, Value: " "},
		{Type: chroma.Name, Value: "world"}}, slices.Collect(tokens))

	var output bytes.Buffer
	assert.NoError(t, quick.Highlight(&output, "hello world", "local-outer", "html", "github"))
	assert.Contains(t, output.String(), "hello")
	assert.Contains(t, output.String(), "world")
	output.Reset()
	assert.NoError(t, quick.Highlight(&output, "<&>\n", "unknown", "html", "github"))
	assert.Contains(t, output.String(), "&lt;&amp;&gt;")
}
