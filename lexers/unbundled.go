//go:build chroma_no_lexers

package lexers

import "github.com/alecthomas/chroma/v3"

func newLexerRegistry() *chroma.LexerRegistry {
	return chroma.NewLexerRegistry()
}
